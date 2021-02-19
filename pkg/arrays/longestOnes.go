package arrays

// LongestOnes 最大连续1的个数 III https://leetcode-cn.com/problems/max-consecutive-ones-iii/
func LongestOnes(arr []int, k int) int {
	offset, zeroCount := 0, 0
	for _, num := range arr {
		if num == 0 {
			zeroCount++
		}
		if zeroCount > k {
			if arr[offset] == 0 {
				zeroCount--
			}
			offset++
		}
	}
	return len(arr) - offset
}
