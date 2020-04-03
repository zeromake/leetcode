package math

// MyPow Pow(x, n) https://leetcode-cn.com/problems/powx-n
func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var ans float64 = 1
	var currentProduct = x
	for i := n; i != 0; i /= 2 {
		if (i % 2) == 1 {
			ans *= currentProduct
		}
		currentProduct *= currentProduct
	}
	return ans
}
