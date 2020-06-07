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
	_ = cr.conn.SetReadDeadline(time.Time{})
	go func() {
		cr.lock()
		n, err := cr.conn.Read(cr.byteBuf[:])
		if n == 1 {
			cr.hasByte = true
		}
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
		} else if err != nil {
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
		// 先处理无 body 的情况
		req = req.WithContext(ctx)
		// 读取请求后开启一个协程读取一个字节来保证检查 net.Conn 的关闭问题
		conn.startBackgroundRead()
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

func handel(w http.ResponseWriter, req *http.Request) {
	go func() {
		select {
		case <-req.Context().Done():

		}
	}()
}
