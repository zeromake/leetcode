package arrays

// FindDuplicate 寻找重复的数 https://leetcode-cn.com/problems/find-the-duplicate-number/
func FindDuplicate(nums []int) int {
	slow, fast := 0, 0
	slow = nums[slow]
	fast = nums[nums[fast]]
	// 对于这个数据以下逻辑会把数组转换为循环链表。
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
