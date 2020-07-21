package links

// DetectCycle 环形链表-II https://leetcode-cn.com/problems/linked-list-cycle-ii/
func DetectCycle(head *ListNode) *ListNode {
	var set = map[*ListNode]struct{}{}
	node := head
	for node != nil {
		if _, ok := set[node]; ok {
			return node
		} else {
			set[node] = struct{}{}
		}
		node = node.Next
	}
	return nil
}
