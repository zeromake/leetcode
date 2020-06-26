package links

// RemoveDuplicateNodes 移除重复节点 https://leetcode-cn.com/problems/remove-duplicate-node-lcci
func RemoveDuplicateNodes(head *ListNode) *ListNode {
	has := make(map[int]struct{})
	node := &ListNode{Next: head}
	for node.Next != nil {
		cur := node.Next
		if _, ok := has[cur.Val]; ok {
			node.Next = node.Next.Next
		} else {
			has[cur.Val] = struct{}{}
			node = cur
		}
	}
	return head
}
