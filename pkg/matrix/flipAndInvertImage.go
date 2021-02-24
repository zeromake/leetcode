package matrix

// FlipAndInvertImage 翻转图像 https://leetcode-cn.com/problems/flipping-an-image
func FlipAndInvertImage(A [][]int) [][]int {
	var n = len(A)
	var m = len(A[0])
	// 必须 +1 把中位数也替换了
	var k = (m+1) / 2
	for i := 0; i < n; i++ {
		row := A[i]
		for j := 0; j < k; j++ {
			row[j], row[m - j-1] = row[m - j-1]^1, row[j]^1
		}
	}
	return A
}
