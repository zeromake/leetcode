package strings

func LetterCombinations(digits string) []string {
	numToLetter := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		if str, ok := numToLetter[digit]; ok {
			pre := res
			res = []string{}
			for j := 0; j < len(str); j++ {
				num := str[j]
				for _, v := range pre {
					res = append(res, v+string(num))
				}
			}
		}
	}
	return res
}
