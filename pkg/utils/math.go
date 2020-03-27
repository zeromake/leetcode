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
