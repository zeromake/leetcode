package arrays

func calculatePoints(boxes []int, dp *[100][100][100]int, left, right, k int) int {
	if left > right {
		return 0
	}
	if dp[left][right][k] != 0 {
		return dp[left][right][k]
	}
	for right > left && boxes[right] == boxes[right-1] {
		right -= 1
		k += 1
	}
	dp[left][right][k] = calculatePoints(boxes, dp, left, right-1, 0) + (k+1)*(k+1)
	for i := left; i < right; i += 1 {
		if boxes[i] == boxes[right] {
			a := calculatePoints(boxes, dp, left, i, k+1) + calculatePoints(boxes, dp, i+1, right-1, 0)
			if a > dp[left][right][k] {
				dp[left][right][k] = a
			}
		}
	}
	return dp[left][right][k]
}

// RemoveBoxes 移除盒子 https://leetcode-cn.com/problems/remove-boxes
func RemoveBoxes(boxes []int) int {
	var dp [100][100][100]int
	return calculatePoints(boxes, &dp, 0, len(boxes)-1, 0)
}
