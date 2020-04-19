package strings

// GetMaxRepetitions 统计重复个数 https://leetcode-cn.com/problems/count-the-repetitions
func GetMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1Count, s2Count, index := 0, 0, 0
	reCall := make(map[int][2]int, 0)
	preLoop, inLoop := [2]int{}, [2]int{}
	for {
		s1Count++
		for i := 0; i < len(s1); i++ {
			c := s1[i]
			if c == s2[index] {
				index++
				if index == len(s2) {
					s2Count++
					index = 0
				}
			}
		}
		if s1Count == n1 {
			return s2Count / n2
		}
		if _, ok := reCall[index]; ok {
			s1CountPrime, s2CountPrime := reCall[index][0], reCall[index][1]
			preLoop = [2]int{s1CountPrime, s2CountPrime}
			inLoop = [2]int{s1Count - s1CountPrime, s2Count - s2CountPrime}
			break
		} else {
			reCall[index] = [2]int{s1Count, s2Count}
		}
	}
	ans := preLoop[1] + (n1-preLoop[0])/inLoop[0]*inLoop[1]
	rest := (n1 - preLoop[0]) % inLoop[0]
	for i := 0; i < rest; i++ {
		for j := 0; j < len(s1); j++ {
			c := s1[j]
			if c == s2[index] {
				index++
				if index == len(s2) {
					ans++
					index = 0
				}
			}
		}
	}
	return ans / n2
}
