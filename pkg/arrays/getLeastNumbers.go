package arrays

import "container/heap"

func GetLeastNumbers(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}
	if k > len(arr) {
		return arr
	}
	h := &MaxHeap{}
	*h = append(*h, arr[:k]...)
	heap.Init(h)
	for i := k; i < len(arr); i++ {
		if arr[i] < (*h)[0] {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	return *h
}

type MaxHeap []int

func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
