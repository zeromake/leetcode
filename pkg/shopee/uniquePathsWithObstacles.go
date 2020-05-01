package shopee

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UniquePathsWithObstacles Shopee的办公室（二）(就是 leetcode 的 不同路径 II) https://www.nowcoder.com/questionTerminal/a71f3bd890734201986cd1e171807d30
func UniquePathsWithObstacles() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	arr := strings.Split(input.Text(), " ")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	n, _ := strconv.Atoi(arr[2])
	// 创建 [x+1][y+1]int 的 dp 并填充为 1
	dp := make([][]int, x+1)
	for i := 0; i <= x; i++ {
		dp[i] = make([]int, y+1)
		for j := 0; j <= y; j++ {
			dp[i][j] = 1
		}
	}
	for i := 0; i < n; i++ {
		input.Scan()
		arr = strings.Split(input.Text(), " ")
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		// 设置路障
		dp[a][b] = 0
	}
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			// 路障处理
			if dp[i][j] == 0 {
				continue
			}
			dp[i][j] = dp[i - 1][j] + dp[i][j-1]
		}
	}
	fmt.Println(dp[x][y])
}
