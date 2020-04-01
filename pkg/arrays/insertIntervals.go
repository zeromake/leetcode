package arrays

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"sort"
)

// InsertIntervals 插入区间 https://leetcode-cn.com/problems/insert-interval/
func InsertIntervals(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	intervals = append(intervals, newInterval)
	sort.Sort(intervalsSort(intervals))
	for _, v := range intervals {
		size := len(result)
		if size == 0 || result[size - 1][1] < v[0] {
			result = append(result, v)
		} else {
			result[size -1][1] = utils.MaxInt(result[size -1][1], v[1])
		}
	}
	return result
}
