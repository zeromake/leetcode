package arrays

// TwoSum 两数相加 https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	sumMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if v, ok := sumMap[target-num]; ok {
			return []int{
				v,
				i,
			}
		}
		sumMap[num] = i
	}
	return nil
}
