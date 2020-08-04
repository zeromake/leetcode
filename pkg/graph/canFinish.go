package graph

// CanFinish 课程表 https://leetcode-cn.com/problems/course-schedule/
func CanFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		index  = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		index[info[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if index[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			index[v]--
			if index[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
