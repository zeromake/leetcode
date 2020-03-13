package arrays

import "sort"

func ThreeSumTarget(nums []int, target int) [][]int {
	var (
		result = make([][]int, 0)
		lens = len(nums)
	)
	if lens < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < lens - 2; i ++ {
		//
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := lens - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{
					nums[i],
					nums[left],
					nums[right],
				})
				for left < right {
					if nums[left] == nums[left + 1] {
						left++
						continue
					}
					if nums[right] == nums[right - 1] {
						right--
						continue
					}
					break
				}
				left ++
				right --
			} else if sum > target {
				right --
			} else {
				left ++
			}
		}
	}
	return result
}

// ThreeSum 三数之和 https://leetcode-cn.com/problems/3sum
func ThreeSum(nums []int) [][]int {
	return ThreeSumTarget(nums, 0)
}
