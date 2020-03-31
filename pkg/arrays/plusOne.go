package arrays

func PlusOne(digits []int) []int {
	mod := 0
	digits[len(digits) - 1] ++
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i] + mod
		mod = digit / 10
		digit %= 10
		digits[i] = digit
	}
	if mod > 0 {
		temp := make([]int, len(digits) + 1)
		temp[0] = mod
		copy(temp[1:], digits)
		return temp
	}
	return digits
}
