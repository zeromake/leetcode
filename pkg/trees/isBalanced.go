package trees

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return utils.MaxInt(height(root.Left), height(root.Right)) + 1
}

// IsBalanced 是否为二叉平衡树 https://leetcode-cn.com/problems/balanced-binary-tree/
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(
		float64(
			height(root.Left)-height(root.Right),
		),
	) < 2 && IsBalanced(root.Left) && IsBalanced(root.Right)
}
