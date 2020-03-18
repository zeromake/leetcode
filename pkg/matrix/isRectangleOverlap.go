package matrix

func IsRectangleOverlap(rec1, rec2 []int) bool {
	// 直接拆分为 x, y 轴来做
	x := !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0])
	y := !(rec1[3] <= rec2[1] || rec2[2] <= rec1[1])
	return x && y
}
