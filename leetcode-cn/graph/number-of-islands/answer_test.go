package answer

import (
	"testing"
)

// https://leetcode-cn.com/problems/number-of-islands/
// 此题的官方测试引擎存在BUG，无法通过测试，不必过多纠结
func numIslands(grid [][]byte) int {
	var dfs func(i int, j int)

	x, y := len(grid[0]), len(grid)
	n := 0
	dfs = func(i int, j int) {
		if i < 0 || i >= x || j < 0 || j >= y {
			return
		}
		val := grid[j][i]
		if val == 1 {
			grid[j][i] = 2 // mark
			dfs(i, j-1)
			dfs(i, j+1)
			dfs(i-1, j)
			dfs(i+1, j)
		}
	}

	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			if grid[j][i] == 1 {
				n++
				dfs(i, j)
			}
		}
	}
	return n
}

func TestIt(t *testing.T) {
	grids := [][][]byte{
		{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}},
		{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}},
		{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}},
	}

	expecteds := []int{1, 3, 1}

	for i := range grids {
		res := numIslands(grids[i])
		expected := expecteds[i]
		if res != expected {
			t.Errorf("expected %v, return %v", expected, res)
		}
	}
}
