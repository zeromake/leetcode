package matrix

var visited map[int]bool
var color map[int]int

func DFS(v int, c int, g [][]int) bool {
	visited[v] = true
	items := g[v]
	color[v] = c
	for _, item := range items {
		if visited[item] != true {
			if !DFS(item, 1-c, g) {
				return false
			}
		} else {
			if color[item] == c {
				return false
			}
		}
	}
	return true
}

// IsBipartite 判断二分图 https://leetcode-cn.com/problems/is-graph-bipartite/
func IsBipartite(graph [][]int) bool {
	visited = make(map[int]bool)
	color = make(map[int]int)
	IsB := true
	for i, _ := range graph {
		if visited[i] != true {
			IsB = DFS(i, 0, graph)
			if IsB != true {
				return IsB
			}
		}
	}
	return IsB
}
