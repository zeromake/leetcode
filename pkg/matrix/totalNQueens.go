package matrix


// TotalNQueens N皇后II https://leetcode-cn.com/problems/n-queens-ii
func TotalNQueens(n int) int {
	res := 0
	dfs3(&res, 0, 0, 0, 0, (1<<n)-1, n)
	return res
}

func dfs3(res *int, count, shu, pie, na, m, n int) {
	if count == n {
		*res ++
		return
	}
	bit := ^(shu | pie | na) & m
	var idx int
	for bit != 0 {
		temp := bit & -bit
		for 1 << idx != temp {idx++}
		dfs3(res, count + 1, shu|temp, (pie|temp)<<1, (na|temp)>>1, m, n)
		bit ^= temp
	}
}
