package arrays

// Intersect 两个数组的交集 II https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
func Intersect(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	nums1Map := make(map[int]int)
	count := 0
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, num1 := range nums1 {
		nums1Map[num1]++
		count++
	}
	for _, num2 := range nums2 {
		if nums1Map[num2] > 0 && count > 0 {
			ret = append(ret, num2)
			nums1Map[num2]--
		}
	}
	return ret
}
