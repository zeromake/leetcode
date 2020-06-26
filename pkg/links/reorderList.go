package links

// ReorderList 重排链表 https://leetcode-cn.com/problems/reorder-list/
func ReorderList(root *ListNode) *ListNode {
	var arr []*ListNode
	for root != nil {
		arr = append(arr, root)
		root = root.Next
	}
	left, right := 0, len(arr)-1
	root = &ListNode{}
	node := root
	for left <= right {
		node.Next = arr[left]
		if left != right {
			node.Next.Next = arr[right]
			node = node.Next.Next
		} else {
			node = node.Next
		}
		left++
		right--
	}
	node.Next = nil
	return root.Next
}
