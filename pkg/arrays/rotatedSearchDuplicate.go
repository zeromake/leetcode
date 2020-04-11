package arrays

// RotatedSearchDuplicate 搜索旋转数组II https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/
func RotatedSearchDuplicate(nums []int, target int) bool {
	var (
		start = 0
		end   = len(nums) - 1
		mid   int
	)
	for start <= end {
		mid = start + (end-start)/2
		v := nums[mid]
		if v == target {
			return true
		}
		if nums[start] == v {
			start++
			continue
		}
		if nums[start] < v {
			// 说明 start ~ mid 之间为升序
			// 检查 target 是否在 start ~ mid 之间
			if v > target && nums[start] <= target {
				// 如果在把 end 直接切到 mid
				end = mid - 1
			} else {
				// 如果不在切换到右边
				start = mid + 1
			}
		} else {
			// 说明 mid ~ end 之间为升序
			// 检查 target 是否在 mid ~ end 之间
			if target > v && target <= nums[end] {
				start = mid + 1
			} else {
				// 不在切换到左边
				end = mid - 1
			}
		}
	}
	return false
}
