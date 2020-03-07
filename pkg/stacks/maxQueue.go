package stacks

// MaxQueue 一个 O(1) 复杂度的队列(可以考虑换成link) https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
type MaxQueue struct {
	queue []int
	deque []int
	size  int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (m *MaxQueue) Max_value() int {
	if len(m.deque) == 0 {
		return -1
	}
	// 直接取出最大值
	return m.deque[0]
}

func (m *MaxQueue) Push_back(value int) {
	if m.size != 0 {
		if value >= m.deque[0] {
			m.deque = m.deque[0:0]
		} else {
			// 尾部小于当前值的已无用，比它小的值只会在该值前被取出
			var i = len(m.deque) - 1
			// 直到大于或者到头部
			for ; i >= 0 && m.deque[i] < value; i-- {
			}
			m.deque = m.deque[:i+1]
		}
	}
	m.deque = append(m.deque, value)
	m.queue = append(m.queue, value)
	m.size++
}

func (m *MaxQueue) Pop_front() int {
	if m.size == 0 {
		return -1
	}
	ans := m.queue[0]
	e := m.deque[0]
	// 检查是否为最大值被剔除
	if ans == e {
		m.deque = m.deque[1:]
	}
	m.queue = m.queue[1:]
	m.size--
	return ans
}
