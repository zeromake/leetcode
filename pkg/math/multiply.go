package math

import (
	"math"
	"strconv"
)

func Multiply(num1, num2 string) string {
	var (
		len1 = len(num1)
		len2 = len(num2)
	)
	sum := 0
	for i := len1 - 1; i >= 0; i -- {
		c1 := int(num1[i] - '0')
		if c1 == 0 {
			continue
		}
		c1 = c1 * int(math.Pow10(len1 - i - 1))
		for j := len2 - 1; j >= 0; j-- {
			c2 := int(num2[j] - '0')
			if c2 == 0 {
				continue
			}
			c2 = c2 * int(math.Pow10(len2 - j - 1))
			sum += c1 * c2
		}
	}
	return strconv.Itoa(sum)
}
