package arrays

import "sort"

// MajorityElement 多数元素 https://leetcode-cn.com/problems/majority-element/
func MajorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func MajorityElement2(nums []int) int {
	var last, count int
	for _, v := range nums {
		if count == 0 {
			last = v
			count = 1
			continue
		}
		if last == v {
			count++
		} else {
			count--
		}
	}
	return last
}
