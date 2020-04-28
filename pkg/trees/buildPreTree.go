package trees

// BuildPreTree 从前序与中序遍历序列构造二叉树 https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func BuildPreTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(preOrder) != len(inOrder) {
		return nil
	}
	inMap := make(map[int]int, len(inOrder))
	for k, v := range inOrder {
		inMap[v] = k
	}
	return buildPreTreeHelper(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1, inMap)
}

func buildPreTreeHelper(preOrder []int, start, end int, intOrder []int, inStart, inEnd int, inMap map[int]int) *TreeNode {
	if start > end || inStart > inEnd {
		return nil
	}
	root := &TreeNode{
		Val: preOrder[start],
	}
	inRoot := 0
	if v, ok := inMap[root.Val]; !ok {
		return nil
	} else {
		inRoot = v
	}
	numLeft := inRoot - inStart
	root.Left = buildPreTreeHelper(preOrder, start+1, end+numLeft, intOrder, inStart, inRoot-1, inMap)
	root.Right = buildPreTreeHelper(preOrder, start+numLeft+1, end, intOrder, inRoot+1, inEnd, inMap)
	return root
}
