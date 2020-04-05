package strings

import "strings"

func IsNumber(s string) bool {
	var isNum = false
	s = strings.TrimSpace(s)
	i := 0
	n := len(s)
	if i < n && ('-' == s[i] || '+' == s[i]) {
		i++
	}
	for i < n && (s[i] >= '0' && s[i] <= '9') {
		isNum = true
		i++
	}
	if i < n && s[i] == '.' {
		i++
	}
	for i < n && (s[i] >= '0' && s[i] <= '9') {
		isNum = true
		i++
	}
	if i < n && (isNum && s[i] == 'e') {
		i++
		isNum = false
		if i < n && ('-' == s[i] || '+' == s[i]) {
			i++
		}
		for i < n && (s[i] >= '0' && s[i] <= '9') {
			i++
			isNum = true
		}
	}
	return i == n && isNum
}

// 编译器状态机方案
func IsNumber2(s string) bool {
	INVALID, SPACE, SIGN, DIGIT, DOT, EXPONENT := 0, 1, 2, 3, 4, 5
	//0invalid,1space,2sign,3digit,4dot,5exponent,6num_inputs
	transitionTable := [][]int{
		{-1, 0, 3, 1, 2, -1},    //0 no input or just spaces
		{-1, 8, -1, 1, 4, 5},    //1 input is digits
		{-1, -1, -1, 4, -1, -1}, //2 no digits in front just dot
		{-1, -1, -1, 1, 2, -1},  //3 sign
		{-1, 8, -1, 4, -1, 5},   //4 digits and dot in front
		{-1, -1, 6, 7, -1, -1},  //5 input 'e' or 'E'
		{-1, -1, -1, 7, -1, -1}, //6 after 'e' input sign
		{-1, 8, -1, 7, -1, -1},  //7 after 'e' input digits
		{-1, 8, -1, -1, -1, -1}} //8 after valid input input space
	state, i, inputtype := 0, 0, 0
	for i < len(s) {
		inputtype = INVALID
		switch s[i] {
		case ' ':
			inputtype = SPACE
			break
		case '-', '+':
			inputtype = SIGN
			break
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			inputtype = DIGIT
			break
		case '.':
			inputtype = DOT
			break
		case 'e', 'E':
			inputtype = EXPONENT
			break
		}
		state = transitionTable[state][inputtype]
		if state == -1 {
			return false
		} else {
			i++
		}
	}
	return state == 1 || state == 4 || state == 7 || state == 8
}
