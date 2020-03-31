package arrays

func SortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	index := partition(nums, left, right)
	QuickSort(nums, left, index-1)
	QuickSort(nums, index+1, right)
}

func partition(nums []int, left int, right int) int {
	i := left
	j := right
	for i != j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}
