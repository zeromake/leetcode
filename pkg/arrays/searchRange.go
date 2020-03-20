package arrays

func SearchRange(nums []int, target int) []int {
	var (
		left  = 0
		right = len(nums) - 1
		mid   = right / 2
		resp = []int{-1, -1}
	)
	if right < 0 {
		return resp
	}
	for left < right {
		v := nums[mid]
		if v == target {
			break
		}
		if v > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
	}
	// 说明没有找到
	if nums[mid] != target {
		return resp
	}

	// 使用中心扩展
	left = mid
	right = mid
	for left > 0 && nums[left-1] == target {
		left--
	}
	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}
	resp[0] = left
	resp[1] = right
	return resp
}

func SearchRange2(nums []int, target int) []int {
	// 确定左边界的二分查找
	begin := leftBinarySearch(nums, target, 0, len(nums)-1)

	// 确定右边界的二分查找
	end := rightBinarySearch(nums, target, 0, len(nums)-1)

	if begin == -1 || end == -1 {
		return []int{-1, -1}
	} else {
		return []int{begin, end}
	}
}

func leftBinarySearch(nums []int, target int, left int, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}

	// 对左侧边界进行检查，防止越界
	if left >= len(nums) || nums[left] != target {
		return -1
	} else {
		return left
	}
}

func rightBinarySearch(nums []int, target int, left int, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}

	// 对左侧边界进行检查，防止越界
	if right < 0 || nums[right] != target {
		return -1
	} else {
		return right
	}
}

func SearchRange3(nums []int, target int) []int {
	// 确定左边界的二分查找
	begin := leftBinarySearch(nums, target, 0, len(nums)-1)
	if begin == -1 {
		return []int{-1, -1}
	}
	right := begin
	for right < len(nums)-1 && nums[right+1] == target {
		right++
	}
	return []int{begin, right}
}

func SearchRange4(nums []int, target int) []int {
	ans := []int{-1, -1}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if ans[0] != -1 {
				ans[1] = i
			} else {
				ans[0] = i
				ans[1] = i
			}
		}
	}
	return ans
}
