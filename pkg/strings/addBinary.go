package strings

func AddBinary(a, b string) string {
	var (
		m, n = len(a), len(b)
		i, j = m - 1, n - 1
	)
	if m < n {
		m, n = n, m
		i, j = j, i
		a, b = b, a
	}
	result := make([]byte, m+1)
	index := m
	var mod uint8
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		ca, cb := a[i]-'0', b[j]-'0'
		sum := ca + cb + mod
		result[index] = (sum % 2) + '0'
		index--
		mod = sum / 2
	}
	for ; i >= 0; i-- {
		c := a[i] - '0'
		sum := c + mod
		result[index] = (sum % 2) + '0'
		index--
		mod = sum / 2
	}
	if mod == 1 {
		result[index] = '1'
		index--
	}
	return string(result[index+1:])
}
