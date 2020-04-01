package arrays

import "github.com/zeromake/leetcode/pkg/utils"

// FindMedianSortedArrays 查找两个数组的中位数 https://leetcode-cn.com/problems/median-of-two-sorted-arrays
func FindMedianSortedArrays(A, B []int) float64 {
	var (
		m      = len(A)
		n      = len(B)
		lens   = m + n
		left   = -1
		right  = -1
		index1 = 0
		index2 = 0
		median = lens / 2
	)
	for i := 0; i <= median; i++ {
		left = right
		if index1 < m && (index2 >= n || A[index1] < B[index2]) {
			right = A[index1]
			index1++
		} else {
			right = B[index2]
			index2++
		}
	}
	if (lens & 1) == 0 {
		return float64(left+right) / 2.0
	}
	return float64(right)
}

func getKth(nums1 []int, start1, end1 int, nums2 []int, start2, end2, k int) int {
	var (
		n = end1 - start1
		m = end2 - start2
	)
	if n > m {
		// 保证 nums2 有元素
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if n == 0 {
		// 发现有数组空了
		return nums2[start2+k-1]
	}
	if k == 1 {
		// 比较到第 1 小的元素直接查找
		return utils.MinInt(nums1[start1], nums2[start2])
	}
	var (
		// 比较的为 k / 2 的元素，作为下标需要减 1
		i = start1 + utils.MinInt(n, k/2) - 1
		j = start2 + utils.MinInt(m, k/2) - 1
	)
	// 判断比第 k 小的元素直接放弃
	if nums1[i] > nums2[j] {
		// 排除已判断出的元素
		k = k - (j - start2 + 1)
		start2 = j + 1
	} else {
		// 排除已判断出的元素
		k = k - (i - start1 + 1)
		start1 = i + 1
	}
	return getKth(
		nums1,
		start1,
		end1,
		nums2,
		start2,
		end2,
		k,
	)
}

// FindMedianSortedArrays2 查找两个数组的中位数
func FindMedianSortedArrays2(A, B []int) float64 {
	var (
		n    = len(A)
		m    = len(B)
		lens = m + n
		left = (lens + 1) / 2
	)
	if (lens & 1) == 0 {
		// 偶数处理
		rigth := (lens + 2) / 2
		return float64(getKth(A, 0, n, B, 0, m, left)+getKth(A, 0, n, B, 0, m, rigth)) / 2.0
	}
	return float64(getKth(A, 0, n, B, 0, m, left))
}

// 和上面的版本完全一样，修改为循环
func getKth2(nums1 []int, nums2 []int, rank int) int {
	var (
		start1 = 0
		start2 = 0
		end1   = len(nums1)
		end2   = len(nums2)
	)
	for rank > 1 {
		n := end1 - start1
		m := end2 - start2
		if n > m {
			nums1, nums2 = nums2, nums1
			n, m = m, n
			start1, start2 = start2, start1
			end1, end2 = end2, end1
		}
		if n == 0 {
			break
		}
		i := start1 + utils.MinInt(rank/2, n) - 1
		j := start2 + utils.MinInt(rank/2, m) - 1
		if nums1[i] > nums2[j] {
			rank = rank - (j - start2 + 1)
			start2 = j + 1
		} else {
			rank = rank - (i - start1 + 1)
			start1 = i + 1
		}
	}
	if end1-start1 == 0 {
		return nums2[start2+rank-1]
	}
	if end2-start2 == 0 {
		return nums1[start1+rank-1]
	}
	return utils.MinInt(nums1[start1], nums2[start2])
}

func FindMedianSortedArrays3(A, B []int) float64 {
	var (
		n    = len(A)
		m    = len(B)
		lens = m + n
		left = (lens + 1) / 2
	)
	if (lens & 1) == 0 {
		// 偶数处理
		right := (lens + 2) / 2
		return float64(getKth2(A, B, left)+getKth2(A, B, right)) / 2.0
	}
	return float64(getKth2(A, B, left))
}
