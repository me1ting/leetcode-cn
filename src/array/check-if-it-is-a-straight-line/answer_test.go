package answer

// https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
// 首先，将整个线上的点平移至经过原点
// 其次对于每个点(x,y)，存在 x1/y1=x2/y2，由于计算机的浮点数的不精确性
// 需要转换为x1*y2=x2*y1，通过整数的相等性来判断
// 事实上，平移的过程也可以放在判断过程中进行，但这里为了直观，参考了题解
//
// 溢出的考虑：题目给出的条件x,y<10000，其乘积在int范围内
func checkStraightLine(coordinates [][]int) bool {
	dX, dY := coordinates[0][0], coordinates[0][1]
	for _, p := range coordinates {
		p[0] -= dX
		p[1] -= dY
	}

	A, B := coordinates[1][0], coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		point := coordinates[i]
		if A*point[1] != B*point[0] {
			return false
		}
	}
	return true
}
