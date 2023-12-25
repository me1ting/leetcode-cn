package main

import "testing"

func rob(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n)

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i := 1; i < n; i++ {
		prev := i - 1
		dp[i][0] = max(dp[prev][0], dp[prev][1])
		dp[i][1] = dp[prev][0] + nums[i]
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestRob(t *testing.T) {
	numsList := [][]int{
		{1, 2, 3, 1},
		{2, 7, 9, 3, 1},
	}

	ansList := []int{4, 12}

	for i := 0; i < len(numsList); i++ {
		nums := numsList[i]
		ans := ansList[i]

		res := rob(nums)
		if res != ans {
			t.Errorf("expected %v, return %v", ans, res)
		}
	}
}
