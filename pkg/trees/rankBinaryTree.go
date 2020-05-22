package trees

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

// 查询该树的节点数量
func size(node *BinaryTree) int {
	if node == nil {
		return 0
	}
	return 1 + size(node.left) + size(node.right)
}

func (p *BinaryTree) Rank(val int) (rank int) {
	if p == nil {
		// 说明没有找到
		rank = 1
	} else if val == p.value {
		// 匹配的排名为左树节点数量 +1
		rank = size(p.left) + 1
	} else if val > p.value {
		// 每次需要去右树查找就加上现有左树的数量
		rank = size(p.left) + p.right.Rank(val) + 1
	} else {
		// 递归到左树查找
		rank = p.left.Rank(val)
	}
	return
}
