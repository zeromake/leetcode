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
	l := make([]int, k)
	i := 0
	var ret [][]int
	for i >= 0 {
		l[i]++
		if l[i] > n-k+i+1 {
			i--
		} else if i == k-1 {
			t := make([]int, k)
			copy(t, l)
			ret = append(ret, t)
		} else {
			i++
			l[i] = l[i-1]
		}
	}
	return ret
}
