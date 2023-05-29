package main

import "testing"

func minPathSum2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)

	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for j := 1; j < m; j++ {
		dp[0] = dp[0] + grid[j][0]
		for i := 1; i < n; i++ {
			dp[i] = min(dp[i-1], dp[i]) + grid[j][i]
		}
	}

	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinPathSum2(t *testing.T) {
	gridList := [][][]int{
		{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		{{1, 2, 3}, {4, 5, 6}},
	}
	rightAnsList := []int{7, 12}

	for i, grid := range gridList {
		rightAns := rightAnsList[i]

		res := minPathSum2(grid)

		if res != rightAns {
			t.Errorf("expected %v, returns %v", rightAns, res)
		}
	}
}
