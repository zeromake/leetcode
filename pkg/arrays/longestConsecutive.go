package arrays

import "sort"

// LongestConsecutive 最长连续序列 https://leetcode-cn.com/problems/longest-consecutive-sequence/
func LongestConsecutive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	sort.Ints(nums)
	count := 1
	max := 1
	for i := 1; i < length; i++ {
		num := nums[i]
		if num == nums[i-1] {
			continue
		} else if num == nums[i-1]+1 {
			count++
		} else {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}

func LongestConsecutive2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := 1
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}
	for num := range set {
		if _, ok := set[num-1]; !ok {
			current := num + 1
			count := 1
			_, ok = set[current]
			for ok {
				count, current = count + 1, current + 1
				_, ok = set[current]
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
