package pa_lou_ti_by_leetcode_solution

// https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode-solution/
// 一个并不简单的简单题
// 考虑最后一步可能跨了一级台阶，也可能跨了两级台阶
// 它意味着爬到第 x 级台阶的方案数是爬到第 x−1 级台阶的方案数和爬到第 x−2级台阶的方案数的和
// 因此类似求斐波那契数
//
// 其不简单之处在于，可以使用矩阵乘法以及数学知识来实现更快的算法
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}