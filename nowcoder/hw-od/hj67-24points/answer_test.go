package hj6724points_test

import (
	"fmt"
	"testing"
)

func main() {
	a, b, c, d := 0, 0, 0, 0
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	fmt.Println(points24([]int{a, b, c, d}))
}

func points24(nums []int) bool {
	visited := make([]bool, len(nums))
	var dfs func(step int, val float64) bool
	dfs = func(step int, val float64) bool {
		if step == 3 {
			if val < 24+0.000001 && val > 24-0.000001 {
				return true
			}
			return false
		}

		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				visited[i] = true
				n := float64(nums[i])

				if dfs(step+1, val+n) {
					return true
				}
				if dfs(step+1, val-n) {
					return true
				}
				if dfs(step+1, val*n) {
					return true
				}
				if dfs(step+1, val/n) {
					return true
				}
				visited[i] = false
			}
		}
		return false
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		visited[i] = true
		if dfs(0, float64(nums[i])) {
			return true
		}
		visited[i] = false
	}
	return false
}

func TestIt(t *testing.T) {
	testcases := [][]int{{9, 6, 4, 1}}
	results := []bool{true}

	for i := 0; i < len(testcases); i++ {
		if points24(testcases[i]) != results[i] {
			t.Fatal("bad answer ", testcases[i], " ", points24(testcases[i]))
		}
	}
}
