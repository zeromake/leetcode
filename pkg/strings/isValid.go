package strings

import "container/list"

// IsValid 检查字符串的括号是否匹配 https://leetcode-cn.com/problems/valid-parentheses
func IsValid(s string) bool {
	stack := list.New()
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '{' || c == '[' {
			stack.PushBack(c + 2)
		} else if c == '(' {
			stack.PushBack(c + 1)
		} else if c == ')' || c == '}' || c == ']' {
			back := stack.Back()
			if back == nil || c != back.Value.(byte) {
				return false
			}
			stack.Remove(back)
		} else {
			return false
		}
	}
	return stack.Front() == nil
}
