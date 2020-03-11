package arrays

// CanThreePartsEqualSum 将数组分成和相等的三个部分 https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
func CanThreePartsEqualSum(arr []int) bool {
	var sum int
	for _, a := range arr {
		sum += a
	}
	if sum % 3 != 0 {
		return false
	}
	num := sum / 3
	left := 0
	leftNum := arr[left]
	right := len(arr) - 1
	rightNum := arr[right]
	for left + 1 < right {
		if leftNum == num && rightNum == num {
			return true
		}
		if leftNum != num {
			left ++
			leftNum += arr[left]
		}
		if rightNum != num {
			right --
			rightNum += arr[right]
		}
	}
	return false
}

