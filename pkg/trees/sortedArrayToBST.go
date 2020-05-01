package trees

// SortedArrayToBST 将有序数组转换为二叉搜索树 https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(nums, 0, len(nums)-1)
}

func sortedArrayToBSTHelper(nums []int, start, end int) *TreeNode {
	if start < 0 || start > end {
		return nil
	}
	mid := (start + end + 1) / 2
	node := &TreeNode{
		Val: nums[mid],
	}
	node.Left = sortedArrayToBSTHelper(nums, start, mid-1)
	node.Right = sortedArrayToBSTHelper(nums, mid+1, end)
	return node
}
