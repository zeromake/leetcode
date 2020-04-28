package trees

// 从后序与中序遍历序列构造二叉树 https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
func BuildPostTree(inOrder []int, postOrder []int) *TreeNode {
	if len(postOrder) == 0 || len(postOrder) != len(inOrder) {
		return nil
	}
	inMap := make(map[int]int, len(inOrder))
	for k, v := range inOrder {
		inMap[v] = k
	}
	end := len(postOrder) - 1
	return buildPostTreeHelper(postOrder, end, 0, len(inOrder)-1, inMap)
}

func buildPostTreeHelper(postOrder []int, end int, inStart, inEnd int, inMap map[int]int) *TreeNode {
	if inStart > inEnd || end < 0 {
		return nil
	}
	root := &TreeNode{
		Val: postOrder[end],
	}
	inRoot := inMap[root.Val]
	// 取得偏移
	numLeft := inEnd - inRoot
	root.Right = buildPostTreeHelper(postOrder, end-1, inRoot+1, inEnd, inMap)
	root.Left = buildPostTreeHelper(postOrder, end-numLeft-1, inStart, inRoot-1, inMap)
	return root
}
