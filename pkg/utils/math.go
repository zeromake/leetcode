package utils

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {

		return a
	}
	return b
}

func Gcd(a, b int) int {
	for {
		if 0 == b {
			break
		}
		a, b = b, a%b
	}
	return a
}

func Factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
