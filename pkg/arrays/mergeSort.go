package arrays

// MergeSort 合并有序数组 https://leetcode-cn.com/problems/merge-sorted-array
func MergeSort(nums1 []int, m int, nums2 []int, n int) {
	result := nums1
	i, j := m-1, n-1
	cur := m + n - 1
	for i >= 0 && j >= 0 {
		num1, num2 := nums1[i], nums2[j]
		if num1 >= num2 {
			result[cur] = num1
			i--

		} else {
			result[cur] = num2
			j--
		}
		cur--
	}
	for j >= 0 {
		result[cur] = nums2[j]
		j--
		cur--
	}
}
