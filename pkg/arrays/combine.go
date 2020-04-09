package arrays

// Combine 组合 https://leetcode-cn.com/problems/combinations
func Combine(n, k int) [][]int {
	result := make([][]int, 0)
	combineDfs(&result, nil, 1, n, k)
	return result
}

func combineDfs(result *[][]int, list []int, start, n, k int) {
	if len(list) == k {
		temp := make([]int, k)
		copy(temp, list)
		*result = append(*result, temp)
		return
	}
	for i := start; i <= n; i++ {
		combineDfs(result, append(list, i), i+1, n, k)
	}
}

func Combine2(n int, k int) [][]int {
	list := make([]int, k)
	i := 0
	var result [][]int
	for i >= 0 {
		// 切换下一个组合
		list[i]++
		// 判断是否已经使用完，使用完回溯到上一个
		if list[i] > n-k+i+1 {
			i--
		} else if i == k-1 {
			// 长度足够插入结果
			t := make([]int, k)
			copy(t, list)
			result = append(result, t)
		} else {
			// 切到一下一个下标
			i++
			list[i] = list[i-1]
		}
	}
	return result
}
