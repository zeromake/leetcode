package links

// LFUCache lfu cache 实现 https://leetcode-cn.com/problems/lfu-cache
type LFUCache struct {
	capacity, size, minFreq int
	freq                    map[int]*DoubleList
	cache                   map[int]*Node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		freq:     make(map[int]*DoubleList),
		cache:    make(map[int]*Node),
	}
}

func (l *LFUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.incFreq(node)
		return node.val
	}
	return -1
}

func (l *LFUCache) Put(key int, value int) {
	if l.capacity == 0 {
		return
	}
	if node, ok := l.cache[key]; ok {
		node.val = value
		l.incFreq(node)
	} else {
		if l.size >= l.capacity {
			// 通过最小的访问数找到需要删除的元素
			node := l.freq[l.minFreq].RemoveLast()
			delete(l.cache, node.key)
			l.size--
		}
		x := &Node{
			key:  key,
			val:  value,
			freq: 1,
		}
		l.cache[key] = x
		if l.freq[1] == nil {
			l.freq[1] = NewDL()
		}
		l.freq[1].Add(x)
		l.minFreq = 1
		l.size++
	}
}

func (l *LFUCache) incFreq(node *Node) {
	freq := node.freq
	l.freq[freq].Remove(node)
	if l.minFreq == freq && l.freq[freq].IsEmpty() {
		l.minFreq++
		delete(l.freq, freq)
	}
	node.freq++
	if l.freq[node.freq] == nil {
		l.freq[node.freq] = NewDL()
	}
	l.freq[node.freq].Add(node)
}

type Node struct {
	prev, next     *Node
	key, val, freq int
}

type DoubleList struct {
	head, tail *Node
	Size       int
}

func (d *DoubleList) Add(node *Node) {
	node.next = d.head.next
	node.prev = d.head

	d.head.next.prev = node
	d.head.next = node
	d.Size++
}

func (d *DoubleList) Remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
	d.Size--
}

func (d *DoubleList) RemoveLast() *Node {
	last := d.tail.prev
	d.Remove(last)
	return last
}

func (d *DoubleList) IsEmpty() bool {
	return d.Size == 0
}

func NewDL() *DoubleList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head
	return &DoubleList{
		head: head,
		tail: tail,
	}
}
