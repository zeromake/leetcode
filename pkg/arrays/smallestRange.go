package arrays

import (
	"container/heap"
	"math"

	"github.com/zeromake/leetcode/pkg/utils"
)

// SmallestRange 最小区间 https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/
func SmallestRange(nums [][]int) []int {
	minx, miny, maxInH := 0, math.MaxInt64, math.MinInt64
	h := new(ItemHeap)
	totalNum := 0
	for i, num := range nums {
		heap.Push(h, &Item{
			value:   num[0],
			numsIdx: i,
			idx:     0,
		})
		maxInH = utils.MaxInt(maxInH, num[0])
		totalNum += len(num)
	}
	for i := 0; i < totalNum; i++ {
		minItem := heap.Pop(h).(*Item)
		if miny-minx > maxInH-minItem.value {
			minx = minItem.value
			miny = maxInH
		}
		if minItem.idx+1 == len(nums[minItem.numsIdx]) {
			break
		}
		heap.Push(h, &Item{
			value:   nums[minItem.numsIdx][minItem.idx+1],
			numsIdx: minItem.numsIdx,
			idx:     minItem.idx + 1,
		})
		maxInH = utils.MaxInt(maxInH, nums[minItem.numsIdx][minItem.idx+1])
	}

	return []int{minx, miny}
}

type Item struct {
	value   int
	numsIdx int
	idx     int
}

type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ItemHeap) Less(i, j int) bool { return h[i].value < h[j].value }

func (h *ItemHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *ItemHeap) Pop() interface{} {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}
