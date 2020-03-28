package math

import "github.com/zeromake/leetcode/pkg/utils"

func CanMeasureWater(x, y, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	// {a: 11, b: 13} -> {a: 0(-11), b: 11(+11)} -> {a: 9(11-2), b: 13(11 + 2)} -> {a: 0, b: 9} -> {a: 7(11-4), b: 13(9+4)}
	// -> {a: 0, b: 7} -> {a: 5(11-6), b: 13(7+6)} -> {a: 0, b: 5} -> {a: 3(11-8), b: 13(5+8)} -> {a: 0, b: 3} -> {a: 1(11-10), b: 13(3+10)}
	return z%utils.Gcd(x, y) == 0
}
