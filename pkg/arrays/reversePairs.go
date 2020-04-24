package arrays

// ReversePairs 数组中的逆序对 https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof
func ReversePairs(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var cp = make([]int, n)
	for i := 0; i < n; i++ {
		cp[i] = nums[i]
	}
	return cntMergeSort(nums, cp, 0, n-1)
}

func cntMergeSort(nums []int, cp []int, start, end int) int {
	if start == end {
		cp[start] = nums[start]
		return 0
	}
	length := (end - start) >> 1
	leftCount := cntMergeSort(cp, nums, start, start+length)
	rightCount := cntMergeSort(cp, nums, start+length+1, end)
	var cnt int
	i, j, copyIndex := start+length, end, end
	for i >= start && j >= start+length+1 {
		if nums[i] > nums[j] {
			cnt += j - start - length
			cp[copyIndex] = nums[i]
			copyIndex--
			i--
		} else {
			cp[copyIndex] = nums[j]
			copyIndex--
			j--
		}
	}
	for i >= start {
		cp[copyIndex] = nums[i]
		copyIndex--
		i--
	}
	for j >= start+length+1 {
		cp[copyIndex] = nums[j]
		copyIndex--
		j--
	}
	return cnt + leftCount + rightCount
}
