package math

func GrayCode(n int) []int {
	result := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		lens := len(result)
		for j := lens - 1; j >= 0; j-- {
			result = append(result, head+result[j])
		}
		head <<= 1
	}
	return result
}
