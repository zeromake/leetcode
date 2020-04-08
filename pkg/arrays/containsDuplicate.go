package arrays

// ContainsDuplicate 检查是否有重复元素 https://leetcode-cn.com/problems/contains-duplicate
func ContainsDuplicate(nums []int) bool {
	cache := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := cache[num]; ok {
			return true
		}
		cache[num] = struct{}{}
	}
	return false
}
