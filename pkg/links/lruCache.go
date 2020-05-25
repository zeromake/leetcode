package links

import "container/list"

// LRUCache lru缓存 https://leetcode-cn.com/problems/lru-cache/
type LRUCache struct {
	cap   int
	data  map[int]*list.Element
	queue *list.List
}

type element struct {
	key   int
	value int
}

func LRUConstructor(cap int) LRUCache {
	c := LRUCache{
		data:  make(map[int]*list.Element),
		cap:   cap,
		queue: list.New(),
	}
	return c
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.data[key]; !ok {
		return -1
	}
	node := c.data[key]
	// 把刚刚访问的元素提取到头部
	c.queue.MoveToFront(node)
	return node.Value.(*element).value
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.data[key]; !ok {
		// 不存在添加到头部并检查是否超过存储空间
		if c.queue.Len() == c.cap {
			b := c.queue.Back()
			v := b.Value.(*element)
			delete(c.data, v.key)
			c.queue.Remove(b)
		}
		node = c.queue.PushFront(&element{key: key, value: value})
		c.data[key] = node
	} else {
		// 存在覆盖并提取到最前
		node.Value.(*element).value = value
		c.queue.MoveToFront(node)
	}
}
