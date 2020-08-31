package matrix

// CanVisitAllRooms 钥匙和房间 https://leetcode-cn.com/problems/keys-and-rooms
func CanVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num := 0
	vis := make([]bool, n)
	var queue []int
	vis[0] = true
	queue = append(queue, 0)
	for i := 0; i < len(queue); i++ {
		x := queue[i]
		num++
		for _, it := range rooms[x] {
			if !vis[it] {
				vis[it] = true
				queue = append(queue, it)
			}
		}
	}
	return num == n
}
