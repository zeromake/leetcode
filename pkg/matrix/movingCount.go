package matrix

func get(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		res += x % 10
	}
	return res
}

func MovingCount(m, n, k int) int {
	if k == 0 {
		return 1
	}
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	count := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || get(i) + get(j) > k {
				continue
			}
			flag := false
			if i - 1 >= 0 {
				flag = flag || dp[i - 1][j]
			}
			if j - 1 >= 0 {
				flag = flag || dp[i][j-1]
			}
			if flag {
				dp[i][j] = true
				count ++
			}
		}
	}
	return count
}

