package answer

import (
	"fmt"
	"testing"
)

func equalPairs(grid [][]int) int {
	n := len(grid[0])
	count := map[string]int{}
	for _, row := range grid {
		key := fmt.Sprint(row)
		count[key]++
	}

	ans := 0

	for i := 0; i < n; i++ {
		col := make([]int, n)
		for j := 0; j < n; j++ {
			col[j] = grid[j][i]
		}
		key := fmt.Sprint(col)
		ans += count[key]
	}
	return ans
}

func TestEqualPairs(t *testing.T) {
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	rightAns := 3

	ans := equalPairs(grid)
	if ans != rightAns {
		t.Errorf("expected %v, but got %v", rightAns, ans)
	}
}
