package matrix

// SetZeroes 矩阵置零 https://leetcode-cn.com/problems/set-matrix-zeroes
func SetZeroes(matrix [][]int) {
	var (
		m, n     = len(matrix), len(matrix[0])
		mdp, ndp = make([]bool, m), make([]bool, n)
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := matrix[i][j]
			if c == 0 {
				mdp[i] = true
				ndp[j] = true
			}
		}
	}
	for i, v := range mdp {
		if v {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j, v := range ndp {
		if v {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}
