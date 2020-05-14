package trees

import (
	"github.com/zeromake/leetcode/pkg/utils"
	"math"
)

// MaxPathSum 最大路径 https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
func MaxPathSum(root *TreeNode) int {
	var maxSum = math.MinInt32
	maxGain(root, &maxSum)
	return maxSum
}

func maxGain(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	left := utils.MaxInt(maxGain(node.Left, maxSum), 0)
	right := utils.MaxInt(maxGain(node.Right, maxSum), 0)
	path := node.Val + left + right
	if path > *maxSum {
		*maxSum = path
	}
	return node.Val + utils.MaxInt(left, right)
}
