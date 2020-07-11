package arrays

import "sort"

type countSmall struct {
	disc  []int
	count []int
}

// CountSmaller 计算右侧小于当前元素的个数 https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
func CountSmaller(nums []int) []int {
	var result []int
	c := &countSmall{
		count: make([]int, len(nums)+5),
	}
	c.discretization(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		id := c.getId(nums[i])
		result = append(result, c.query(id-1))
		c.update(id)
	}
	lens := len(result)
	for i := 0; i < lens/2; i++ {
		result[i], result[lens-1-i] = result[lens-1-i], result[i]
	}
	return result
}

func lowBit(x int) int {
	return x & (-x)
}

func (c *countSmall) update(pos int) {
	for pos < len(c.count) {
		c.count[pos]++
		pos += lowBit(pos)
	}
}

func (c *countSmall) query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += c.count[pos]
		pos -= lowBit(pos)
	}
	return ret
}

func (c *countSmall) discretization(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	c.disc = make([]int, 0, len(nums))
	for num := range set {
		c.disc = append(c.disc, num)
	}
	sort.Ints(c.disc)
}

func (c *countSmall) getId(x int) int {
	return sort.SearchInts(c.disc, x) + 1
}
