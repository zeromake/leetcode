package trees

import "github.com/zeromake/leetcode/pkg/utils"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

// DiameterOfBinaryTree 查找二叉树最长路径 https://leetcode-cn.com/problems/diameter-of-binary-tree
func DiameterOfBinaryTree(root *treeNode) int {
	var (
		count = 0
	)
	depth(root, &count)
	return count
}

func depth(node *treeNode, max *int) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left, max)
	right := depth(node.Right, max)
	*max = utils.MaxInt(*max, left+right)
	return utils.MaxInt(left, right) + 1
}
