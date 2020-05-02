package trees

type ListNode struct {
	Val  int
	Next *ListNode
}

// SortedListToBST 有序链表转二叉搜索树 https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
func SortedListToBST(head *ListNode) *TreeNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return sortedArrayToBSTHelper(arr, 0, len(arr)-1)
}
