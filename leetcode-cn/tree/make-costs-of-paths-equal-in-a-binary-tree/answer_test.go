package makecostsofpathsequalinabinarytree_test

// https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/
// 思路：从索引最大的非叶子节点往左遍历，使得其两个子节点平衡（即相等，这里的实现并没有使其相等，而是增加minInc）
//      然后将平衡后的子路径大小记录在节点上
func minIncrements(n int, cost []int) int {
	minInc := 0

	// firstLeaf := n / 2
	for i := n/2 - 1; i >= 0; i-- {
		left := i*2 + 1

		m, n := cost[left], cost[left+1]

		if m > n {
			minInc += m - n
			cost[i] += m
		} else {
			minInc += n - m
			cost[i] += n
		}
	}

	return minInc
}
