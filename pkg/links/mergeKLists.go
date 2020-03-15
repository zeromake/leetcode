package links

// MergeKLists 合并多个有序链表 https://leetcode-cn.com/problems/merge-k-sorted-lists
func MergeKLists(lists []*ListNode) *ListNode {
	var (
		lens     = len(lists)
		interval = 1
	)
	for interval < lens {
		// 长度减去 interval 处理非偶数长度
		for i := 0; i < lens-interval; i += interval * 2 {
			// 合并步进长度 interval
			lists[i] = MergeTwoLists(lists[i], lists[i+interval])
		}
		// 调节步进长度保证下次的合并是合并后的
		interval *= 2
	}
	if lens > 0 {
		return lists[0]
	}
	return nil
}
