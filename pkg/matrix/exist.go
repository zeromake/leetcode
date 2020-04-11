package matrix

func Exist(board [][]byte, word string) bool {
	var (
		m, n = len(board), len(board[0])
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := board[i][j]
			if c == word[0] {
				if existDfs(board, word, m, n, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

func existDfs(board [][]byte, word string, m, n, x, y, count int) bool {
	if count == len(word) {
		return true
	}
	if x < 0 || x >= m || y < 0 || y >= n {
		return false
	}
	if board[x][y] != word[count] {
		return false
	}
	curr := board[x][y]
	board[x][y] = ' '
	flag := false
	for _, dir := range dirs {
		i, j := x+dir[0], y+dir[1]
		flag = flag || existDfs(board, word, m, n, i, j, count+1)
	}
	board[x][y] = curr
	return flag
}

func Exist2(nums [][]byte, word string) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			if helper(nums, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func helper(nums [][]byte, i, j int, word string, index int) bool {
	// 1.递归退出条件，index== 单词长度  true
	if len(word) == index {
		return true
	}
	// 2.递归退出条件， 越界
	if i < 0 || i == len(nums) || j < 0 || j == len(nums[0]) {
		return false
	}
	// 3.递归退出条件, word当前值与nums中不匹配
	if nums[i][j] != word[index] {
		return false
	}
	// 进入递归逻辑
	curr := nums[i][j]
	nums[i][j] = byte(' ') // 防止重复
	ret := helper(nums, i, j+1, word, index+1) ||
		helper(nums, i, j-1, word, index+1) ||
		helper(nums, i+1, j, word, index+1) ||
		helper(nums, i-1, j, word, index+1)
	nums[i][j] = curr // 回溯
	return ret
}
