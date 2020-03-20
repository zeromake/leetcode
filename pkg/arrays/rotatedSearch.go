package arrays

func RotatedSearch(nums []int, target int) int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   = right / 2
	)
	for left <= right {
		v := nums[mid]
		if v == target {
			return mid
		}
		if nums[left] <= v {
			// 说明 left ~ mid 之间为升序
			// 检查 target 是否在 left ~ mid 之间
			if target >= nums[left] && target < v {
				// 如果在把 right 直接切到 mid
				right = mid - 1
			} else {
				// 如果不在切换到右边
				left = mid + 1
			}
		} else {
			// 说明 mid ~ right 之间为升序
			// 检查 target 是否在 mid ~ right 之间
			if target > v && target <= nums[right] {
				left = mid + 1
			} else {
				// 不在切换到左边
				right = mid - 1
			}
		}
		mid = left + (right-left)/2
	}
	return -1
}
