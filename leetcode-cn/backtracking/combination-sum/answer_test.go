package combinationsum

// https://leetcode.cn/problems/combination-sum
// 使用递归求解，思路类似问题subsets的解法2
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	heap := []int{}

	var dfs func(n int, t int)
	dfs = func(n int, t int) {
		if t == 0 {
			res = append(res, append([]int{}, heap...))
			return
		}
		if n == len(candidates) {
			return
		}
		num := candidates[n]
		i := 0
		for ; i*num <= t; i++ {
			if i > 0 {
				heap = append(heap, num)
			}

			dfs(n+1, t-i*num)
		}
		heap = heap[:len(heap)-i+1]
	}

	dfs(0, target)
	return res
}
