package minimum_number_of_arrows_to_burst_balloons

import "sort"

// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons
// 算法本身并不复杂，只是需要转换思维
// https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/solution/yong-zui-shao-shu-liang-de-jian-yin-bao-qi-qiu-1-2/
// 我们将所有气球看作线段，将所有线段按照右边距排序：
// 1. 每个线段必然存在一个箭头穿过它，让当前线段的最右侧作为射箭的位置，此时对当前线段来讲，这个位置是最佳位置，
// 因为可以穿过尽可能多的线段
// 2. 贪婪法可以解决这个问题：首先，我们可以自然得出，对于与当前线段无交叉的后续线段来讲，当前线段的箭头的位置对
//    它们是没有影响的；其次，剩余越少的线段集合需要的箭数只少不多。因此局部最优等同于全局最优。
// 3. 存在部分小线段位于当前线段的箭头所射入的线段群内，因此要按照上面那样的想法实现代码需要遍历移除被射入的线段
//    可以优化为：遇到不相交的线段就视为射下一个箭，而那些被小线段遮挡的长线段可以被下一箭所射中
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(Points(points))
	pos := points[0][1]
	count := 1
	for _, point := range points[1:] {
		if pos < point[0] {
			pos = point[1]
			count++
		}
	}
	return count
}
type Points [][]int

func (a Points) Len() int           { return len(a) }
func (a Points) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Points) Less(i, j int) bool { return a[i][1] < a[j][1] }