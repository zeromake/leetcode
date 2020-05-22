## 二叉搜索树的排名查找
> 给定一个二叉搜索树和一个值，查找该值在这个树中排名，该值有可能不存在与树中

```go
type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

// 查询该树的节点数量
func size(node *BinaryTree) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

func (p *BinaryTree) Rank(val int) (rank int) {
	if p == nil {
		// 说明没有找到
		rank = 1
	} else if val == p.value {
		// 匹配的排名为左树节点数量 +1
		rank = size(p.left) + 1
	} else if val > p.value {
		// 每次需要去右树查找就加上现有左树的数量
		rank = size(p.left) + p.right.Rank(val) + 1
	} else {
		// 递归到左树查找
		rank = p.left.Rank(val)
	}
	return
}
```

## http 代理时的断开信号检查
> 以代理 http1.1 为准，如何在一个请求读取后继续监听连接并且传递到代理方。

参考 [connReader](https://github.com/golang/go/blob/release-branch.go1.14/src/net/http/server.go#L642-L798)

```go
package proxy

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"sync"
	"time"
)

type connReader struct {
	mu        sync.Mutex
	cond      *sync.Cond
	conn      net.Conn
	hasByte   bool
	byteBuf   [1]byte
	inRead    bool
	cancelCtx func()
}

// 包装 Read 支持多读取的一个字节
func (cr *connReader) Read(p []byte) (n int, err error) {
	cr.lock()
	if cr.inRead {
		cr.unlock()
		panic("invalid concurrent Body.Read call")
	}
	if len(p) == 0 {
		cr.unlock()
		return 0, nil
	}
	if cr.hasByte {
		p[0] = cr.byteBuf[0]
		cr.hasByte = false
		cr.unlock()
		return 1, nil
	}
	cr.inRead = true
	cr.unlock()
	n, err = cr.conn.Read(p)

	cr.lock()
	cr.inRead = false
	if err != nil {
		cr.handleReadError(err)
	}
	cr.unlock()

	cr.cond.Broadcast()
	return n, err
}

func (cr *connReader) handleReadError(_ error) {
    // 简化只调用一个 cancel
	cr.cancelCtx()
}
func (cr *connReader) lock() {
	cr.mu.Lock()
	if cr.cond == nil {
		cr.cond = sync.NewCond(&cr.mu)
	}
}

func (cr *connReader) unlock() { cr.mu.Unlock() }

func (cr *connReader) startBackgroundRead() {
	cr.lock()
	defer cr.unlock()
	if cr.inRead {
		panic("invalid concurrent Body.Read call")
	}
	if cr.hasByte {
		return
	}
	cr.inRead = true
    // 设置不超时
	_ = cr.conn.SetReadDeadline(time.Time{})
	go func() {
		cr.lock()
        // 以读取一个字节来检查连接情况
		n, err := cr.conn.Read(cr.byteBuf[:])
		if n == 1 {
			cr.hasByte = true
		}
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
		} else if err != nil {
            // EOF 错误会到这里
			cr.handleReadError(err)
		}
		cr.inRead = false

		cr.unlock()
		cr.cond.Broadcast()
	}()
}

func Proxy(origin net.Conn, target net.Conn) error {
	ctx, cancel := context.WithCancel(context.Background())
	conn := &connReader{
		conn:      origin,
		byteBuf:   [1]byte{},
		cancelCtx: cancel,
	}
	reader := bufio.NewReader(conn)
	rb := bufio.NewReader(target)
	for {
		ctx, cancel = context.WithCancel(context.Background())
		conn.cancelCtx = cancel
		req, err := http.ReadRequest(reader)
		if err != nil {
			return err
		}
		// 先不考虑无 body 的情况，如果需要考虑body需要等待body读取完成再调用startBackgroundRead
		req = req.WithContext(ctx)
		// 读取请求后开启一个协程读取一个字节来保证检查 net.Conn 的关闭问题
		conn.startBackgroundRead()
        // 应该修改为 http.Client 去发起请求才能支持 ctx 的请求取消能力
		err = req.Write(target)
		if err != nil {
			return err
		}
		resp, err := http.ReadResponse(rb, req)
		if err != nil {
			return err
		}
		err = resp.Write(origin)
		if err != nil {
			return err
		}
	}
}
```
## 如何检查一个 `api` 访问性能和稳定性
> 响应时间怎么做监控，如何查看少量波动。

参考 `Prometheus` 的 `Histogram` 和 `Summary`。

**Histogram**
根据指定的 bucket 时间段(0.005s, 0.01s 之类的)，划分 0.01s,0.1s 内的响应数量，这样可以很清楚的看出响应时间分布。
```
# HELP http_request_duration_seconds Histogram of lantencies for HTTP requests
# TYPE http_request_duration_seconds histogram
http_request_duration_seconds_bucket{le="0.005"} 0
http_request_duration_seconds_bucket{le="0.01"} 0
http_request_duration_seconds_bucket{le="0.025"} 0
http_request_duration_seconds_bucket{le="0.05"} 0
http_request_duration_seconds_bucket{le="0.1"} 3
http_request_duration_seconds_bucket{le="0.25"} 3
http_request_duration_seconds_bucket{le="0.5"} 5
http_request_duration_seconds_bucket{le="1"} 8
http_request_duration_seconds_bucket{le="2.5"} 8
http_request_duration_seconds_bucket{le="5"} 8
http_request_duration_seconds_bucket{le="10"} 8
http_request_duration_seconds_bucket{le="+Inf"} 8
http_request_duration_seconds_sum 3.238809838
http_request_duration_seconds_count 8
```

**Summary**
一般使用 `Histogram` 就够了，`Summary` 根据设置的分位数进行平均值计算。

```
$ curl http://127.0.0.1:8080/metrics | grep http_request_summary
# HELP http_request_summary_seconds Summary of lantencies for HTTP requests
# TYPE http_request_summary_seconds summary
http_request_summary_seconds{quantile="0.5"} 0.31810446
http_request_summary_seconds{quantile="0.9"} 0.887116164
http_request_summary_seconds{quantile="0.99"} 0.887116164
http_request_summary_seconds_sum 3.2388269649999994
http_request_summary_seconds_count 8
```
## 参考

- [如何利用Prometheus监控你的应用](https://www.cnblogs.com/YaoDD/p/11391316.html)
