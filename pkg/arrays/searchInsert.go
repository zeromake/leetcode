package arrays

func SearchInsert(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   = left + (right-left)/2
	)
	// 由于是插入位置到达边界需要再走一次
	for left <= right {
		// 二分查找
		v := nums[mid]
		if v == target {
			return mid
		} else if v > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
	}
	return mid
}
