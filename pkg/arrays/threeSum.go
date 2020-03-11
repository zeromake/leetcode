package arrays

import "sort"

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func ThreeSumTarget(nums []int, target int) [][]int {
	var (
		result = make([][]int, 0)
		lens = len(nums)
	)
	sort.Sort(IntSlice(nums))
	for i := 0; i < lens - 1; i ++ {
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

func ThreeSum(nums []int) [][]int {
	return ThreeSumTarget(nums, 0)
}
