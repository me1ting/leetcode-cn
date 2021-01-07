package answer

import (
	"testing"
)

// https://leetcode-cn.com/problems/number-of-provinces
// 官方题解：https://leetcode-cn.com/problems/number-of-provinces/solution/sheng-fen-shu-liang-by-leetcode-solution-eyk0/
// 求解连通分量的标准题
func findCircleNum(isConnected [][]int) int {
	marks := make([]bool, len(isConnected))
	n := 0

	var dfs func(int)

	//这里为了方便，使用递归式的dfs
	dfs = func(i int) {
		marks[i] = true
		for j := 0; j < len(marks); j++ {
			if j != i && isConnected[i][j] == 1 && !marks[j] {
				dfs(j)
			}
		}
	}

	for i := 0; i < len(marks); i++ {
		if !marks[i] {
			n++
			dfs(i)
		}
	}

	return n
}

func TestIt(t *testing.T) {
	isConnected := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	if findCircleNum(isConnected) != 3 {
		t.Errorf("error")
	}
}
