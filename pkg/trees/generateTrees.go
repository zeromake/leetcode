package trees

// GenerateTrees https://leetcode-cn.com/problems/unique-binary-search-trees-ii
func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTrees(1, n)
}

func generateTrees(start, end int) []*TreeNode {
	var trees []*TreeNode
	if start > end {
		trees = append(trees, nil)
		return trees
	}
	for i := start; i <= end; i++ {
		lefts := generateTrees(start, i-1)
		rights := generateTrees(i+1, end)
		for _, l := range lefts {
			for _, r := range rights {
				cur := &TreeNode{Val: i}
				cur.Left = l
				cur.Right = r
				trees = append(trees, cur)
			}
		}
	}
	return trees
}
