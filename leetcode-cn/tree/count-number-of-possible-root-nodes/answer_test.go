package countnumberofpossiblerootnodes_test

// https://leetcode.cn/problems/count-number-of-possible-root-nodes/
// 题解：https://leetcode.cn/problems/count-number-of-possible-root-nodes/solutions/2654312/tong-ji-ke-neng-de-shu-gen-shu-mu-by-lee-3gzg/
// 思路：
// 通过一次dfs，统计以节点0作为树根时，正确的猜测的边数count
// (core)当根节点从0转换为0时，只有0-1边变化为1-0边，其它边的方向并未改变，此时的count只需要计入该边变化对count产生的影响
// 因此，再通过一次dfs，可以统计所有count满足count>=k的次数
func rootCount(edges [][]int, guesses [][]int, k int) int {
	n := len(edges) + 1
	g := make([][]int, n)
	t := map[int64]bool{}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for _, guess := range guesses {
		a, b := guess[0], guess[1]
		t[hash(a, b)] = true
	}

	count := 0
	var dfs func(int, int)
	dfs = func(a int, from int) {
		for _, b := range g[a] {
			if b == from {
				continue
			}
			if t[hash(a, b)] {
				count++
			}
			dfs(b, a)
		}
	}
	dfs(0, -1)

	res := 0
	var reDfs func(int, int, int)
	reDfs = func(a int, from int, cnt int) {
		if t[hash(from, a)] {
			cnt--
		}
		if t[hash(a, from)] {
			cnt++
		}

		if cnt >= k {
			res++
		}
		for _, b := range g[a] {
			if b == from {
				continue
			}
			reDfs(b, a, cnt)
		}
	}
	reDfs(0, -1, count)

	return res
}

func hash(a, b int) int64 {
	return int64(a)<<17 | int64(b)
}
