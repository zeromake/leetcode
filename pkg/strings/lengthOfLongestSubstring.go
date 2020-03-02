package strings

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		// 转换为 []byte 阻止 rune 转换
		bytes = []byte(s)
		// 记录上次出现字符的位置使用 []int 加速
		cache = make([]int, 128)
		// 结果和当前滑动窗口大小
		max, num int
		// 滑动窗口左游标
		left int
		// 字符串长度
		lens = len(bytes)
	)
	for right := 0; left < lens && right < lens; right++ {
		b := bytes[right]
		// 发现有相同字符把左游标移动到上次出现的下一个
		if n := cache[b]; n > left {
			left = n
		}
		// 判断当前滑动窗口大小
		num = right - left + 1
		if max < num {
			max = num
		}
		// 存放当前字符的下一个 index
		cache[b] = right + 1
	}
	return max
}
