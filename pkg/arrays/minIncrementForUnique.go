package arrays

// https://leetcode-cn.com/problems/minimum-increment-to-make-array-unique/
func MinIncrementForUnique(nums []int) int {
	pos := make(map[int]int, 0)
	move := 0
	for _, num := range nums {
		m := findPos(pos, num)
		move += m - num
	}
	return move
}

func findPos(pos map[int]int, num int) int {
	if m, ok := pos[num]; ok {
		m = findPos(pos, m + 1)
		pos[num] = m
		return m
	} else {
		pos[num] = num
		return num
	}
}

func MinIncrementForUnique2(nums []int) int {
	var (
		pos = make(map[int]int, 0)
		min, max = nums[0], nums[0]
	)
	for _, num := range nums {
		pos[num]++
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	move := 0
	for min < max {
		if pos[min] > 1 {
			dup := pos[min] - 1
			move += dup
			pos[min] = 1
			pos[min + 1] += dup
		}
		min ++
	}
	dup := pos[max] - 1
	move += dup * (dup + 1) / 2
	return move
}
