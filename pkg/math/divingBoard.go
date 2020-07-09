package math

// DivingBoard 跳水板 https://leetcode-cn.com/problems/diving-board-lcci/
func DivingBoard(shorter int, longer int, k int) []int {
	switch k {
	case 0:
		return nil
	case 1:
		return []int{shorter, longer}
	}
	if shorter == longer {
		return []int{k * shorter}
	}
	result := make([]int, k+1)
	for i := 0; i <= k; i++ {
		result[i] = i*longer + (k-i)*shorter
	}
	return result
}
