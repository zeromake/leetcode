package matrix

// SolveNQueens N皇后 https://leetcode-cn.com/problems/n-queens
func SolveNQueens(n int) [][]string {
	var result [][]string
	var matrix = make([][]byte, n)
	c := make([]byte, n)
	for i := 0; i < n; i++ {
		c[i] = '.'
	}
	matrix[0] = c
	for i := 1; i < n; i++ {
		matrix[i] = make([]byte, n)
		copy(matrix[i], c)
	}
	dfs(&result, matrix, 0, n, 0)
	return result
}

func hasQueens(x, y, n int, matrix [][]byte) bool {
	for i := 0; i < n; i++ {
		if matrix[i][y] == 'Q' {
			return true
		}
	}
	for i, j := x-1, y+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if matrix[i][j] == 'Q' {
			return true
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if matrix[i][j] == 'Q' {
			return true
		}
	}
	return false
}

func dfs(result *[][]string, matrix [][]byte, x, n, num int) {
	if num == n {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = string(matrix[i])
		}
		*result = append(*result, temp)
		return
	}

	for y := 0; y < n; y++ {
		if hasQueens(x, y, n, matrix) {
			continue
		}
		matrix[x][y] = 'Q'
		dfs(result, matrix, x+1, n, num+1)
		matrix[x][y] = '.'
	}
}

func SolveNQueens2(n int) [][]string {
	res := make([][]string, 0)
	brr := make([]byte, n)
	for i := 0; i < n; i++ {
		brr[i] = '.'
	}
	dfs2(&res, make([]string, 0), brr, 0, 0, 0, (1<<n)-1, n)
	return res
}

func dfs2(res *[][]string, arr []string, brr []byte, shu, pie, na, m, n int) {
	if len(arr) == n {
		*res = append(*res, myCopy(arr, n))
		return
	}
	bit := ^(shu | pie | na) & m
	var idx int
	for bit != 0 {
		temp := bit & -bit
		for 1<<idx != temp {
			idx++
		}
		brr[idx] = 'Q'
		arr = append(arr, string(brr))
		brr[idx] = '.'
		dfs2(res, arr, brr, shu|temp, (pie|temp)<<1, (na|temp)>>1, m, n)
		bit ^= temp
		arr = arr[:len(arr)-1]
	}
}

func myCopy(arr []string, n int) []string {
	brr := make([]string, n)
	copy(brr, arr)
	return brr
}
