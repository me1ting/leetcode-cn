package findeventualsafestates

// https://leetcode-cn.com/problems/find-eventual-safe-states/
// 独立思考：
// 1.图可能存在环，环上的所有点和能够到达环的所有点都是非安全起始点。
// 2.终点的入度为0
// 但是无法独立解决该问题，需要参考题解。
//
// https://leetcode-cn.com/problems/find-eventual-safe-states/solution/zhao-dao-zui-zhong-de-an-quan-zhuang-tai-yzfz/
// 题解：使用深度优先搜索和三色标记法
// 从代码实现来讲，很巧妙和精炼，通过 递归+记忆化 隐式的实现了深度优先搜索

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	color := make([]int, n) //0: not visited;1: on cycle or on stack;2: unsafe
	var safe func(v int) bool
	safe = func(v int) bool {
		if color[v] > 0 {
			return color[v] == 2
		}
		color[v] = 1
		for _, w := range graph[v] {
			if !safe(w) {
				return false
			}
		}
		color[v] = 2
		return true
	}

	ans := []int{}
	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
