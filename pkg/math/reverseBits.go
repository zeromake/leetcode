package math

// ReverseBits 颠倒二进制位 https://leetcode-cn.com/problems/reverse-bits
func ReverseBits(num uint32) uint32 {
	var ans uint32 = 0
	var i uint32 = 32
	for i > 0 {
		ans <<= 1
		ans += num & 1
		num >>= 1
		i--
	}
	return ans
}
