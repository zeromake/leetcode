package math

import "math"

const min = math.MinInt32 >> 1

func Divide(dividend, divisor int) int {
	// 只有不是同正负符号
	flag := (dividend ^ divisor) < 0
	result := 0
	if dividend > 0 {
		dividend = 0 - dividend
	}
	if divisor > 0 {
		divisor = 0 - divisor
	}
	for dividend <= divisor {
		tempResult := -1
		tempDivisor := divisor
		// 核心算法，不停的对除数乘 2, 直到大于被除数，-1 也在这其中乘 2 最后的结果就是商，在乘 2 的过程中可能会有余数
		for dividend <= (tempDivisor + tempDivisor) {
			if tempDivisor <= min {
				break
			}
			// 记录除数数量
			tempResult = tempResult + tempResult
			// 对除数一直 * 2
			tempDivisor = tempDivisor + tempDivisor
		}
		// 被除数减去
		dividend = dividend - tempDivisor
		// 加上减去的除数数量
		result += tempResult
	}
	if !flag {
		if result <= math.MinInt32 {
			return math.MaxInt32
		}
		result = 0 - result
	}
	return result
}

//func Divide2(dividend, divisor int) int {
//	// 只有不是同正负符号
//	flag := (dividend ^ divisor) < 0
//	result := 0
//	if dividend > 0 {
//		dividend = 0 - dividend
//	}
//	if divisor > 0 {
//		divisor = 0 - divisor
//	}
//	tempDivisor := divisor
//	// 核心算法，不停的减去除数
//	for dividend <= tempDivisor {
//		// 记录除数数量
//		result -= 1
//		if tempDivisor <= math.MinInt32 {
//			break
//		}
//		// 对除数一直 * 2
//		tempDivisor += divisor
//	}
//	if !flag {
//		if result <= math.MinInt32 {
//			return math.MaxInt32
//		}
//		result = 0 - result
//	}
//	return result
//}

