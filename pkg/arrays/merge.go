package arrays

// Merge 合并两个有序数组 https://leetcode-cn.com/problems/sorted-merge-lcci/
func Merge(A []int, m int, B []int, n int) {
	cur := m + n - 1
	index1 := m - 1
	index2 := n - 1
	for index1 >= 0 && index2 >= 0 {
		x := A[index1]
		y := B[index2]
		if x < y {
			A[cur] = y
			index2--
		} else {
			A[cur] = x
			index1--
		}
		cur--
	}
	if index2 >= 0 {
		copy(A[:index2+1], B[:index2+1])
	}
}
