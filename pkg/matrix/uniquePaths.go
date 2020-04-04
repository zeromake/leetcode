package matrix

import "github.com/zeromake/leetcode/pkg/utils"


func UniquePaths(m, n int) int {
	var cache = make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				cache[i][j] = 1
			} else {
				cache[i][j] = cache[i - 1][j] + cache[i][j - 1]
			}
		}
	}
	return cache[m-1][n-1]
}

func UniquePaths2(m, n int) int {
	if m > n {
		m, n = n, m
	}
	temp, result := utils.Factorial(m - 1), 1
	// 通过手动控制把 factorial(n - 1) 去除防止溢出
	for i := n; i <= m + n - 2; i++ {
		result *= i
	}
	result /= temp
	return result
}
