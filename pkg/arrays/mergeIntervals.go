package arrays

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"sort"
)

type intervalsSort [][]int

func (ii intervalsSort) Len() int {
	return len(ii)
}

func (ii intervalsSort) Less(i, j int) bool {
	return ii[i][0] < ii[j][0]
}

func (ii intervalsSort) Swap(i, j int) {
	ii[i], ii[j] = ii[j], ii[i]
}

// MergeIntervals 合并区间 https://leetcode-cn.com/problems/merge-intervals
func MergeIntervals(intervals [][]int) [][]int {
	var result [][]int
	if len(intervals) == 0 {
		return result
	}
	sort.Sort(intervalsSort(intervals))
	var left = intervals[0][0]
	var right = intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		n, m := intervals[i][0], intervals[i][1]
		if !(n > right || m < left) {
			right = utils.MaxInt(right, m)
			left = utils.MinInt(left, n)
		} else {
			result = append(result, []int{left, right})
			left = n
			right = m
		}
	}
	result = append(result, []int{left, right})
	return result
}

func MergeIntervals2(intervals [][]int) [][]int {
	var result [][]int
	sort.Sort(intervalsSort(intervals))
	for _, v := range intervals {
		size := len(result)
		if size == 0 || result[size-1][1] < v[0] {
			result = append(result, v)
		} else {
			result[size-1][1] = utils.MaxInt(result[size-1][1], v[1])
		}
	}
	return result
}
