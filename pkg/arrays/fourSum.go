package arrays

import "sort"

func FourSum(nums []int, target int) [][]int {
	var (
		result = make([][]int, 0)
		lens = len(nums)
	)
	if lens < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < lens - 3; i ++ {
		//
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < lens - 2; j++ {
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := lens - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{
						nums[i],
						nums[j],
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

	}
	return result
}
