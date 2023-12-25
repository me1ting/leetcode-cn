package answer

// https://leetcode-cn.com/problems/fair-candy-swap/
// 这本质是一道数学题：
// 因为需要交换一对糖果x,y，那么 sum_A-x+y=sum_B+x-y
// x-y = (sum_A-sum_B)/2
// y = x-(sum_A-sum_B)/2
func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	aRecords, bRecords := map[int]bool{}, map[int]bool{}

	for _, n := range A {
		sumA += n
		aRecords[n] = true
	}

	for _, n := range B {
		sumB += n
		bRecords[n] = true
	}

	c := (sumA - sumB) / 2

	for _, n := range A {
		y := n - c
		if bRecords[y] == true {
			return []int{n, y}
		}
	}
	return nil
}
