package math

// WaysToChange 硬币 https://leetcode-cn.com/problems/coin-lcci/
func WaysToChange(n int) int {
	ans := 0
	const mod = 1000000007
	for i := 0; i*25 <= n; i++ {
		rest := n - i*25
		a := rest / 10
		b := rest % 10 / 5
		ans = (ans + (a+1)*(a+b+1)%mod) % mod
	}
	return ans
}
