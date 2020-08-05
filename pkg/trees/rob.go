package trees

import "github.com/zeromake/leetcode/pkg/utils"

// Rob 打家劫舍 III https://leetcode-cn.com/problems/house-robber-iii/
func Rob(root *TreeNode) int {
	val := robDfs(root)
	return utils.MaxInt(val[0], val[1])
}

func robDfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := robDfs(node.Left), robDfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := utils.MaxInt(l[0], l[1]) + utils.MaxInt(r[0], r[1])
	return []int{selected, notSelected}
}
