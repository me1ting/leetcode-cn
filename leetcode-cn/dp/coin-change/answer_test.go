package coinchange_test

import (
	"sort"
	"testing"
)

// https://leetcode.cn/problems/coin-change/
// 思路：
// 使用dp，和所有dp问题一样，核心问题是如何找到状态方程。
// 对于每个位置i，dp[i] = min(dp[c]+dp[i-c])，要求dp[c]和dp[i-c]有效
// 因此算法的复杂度为O(len(coins)*n) 即 O(n)
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Ints(coins)
	if coins[0] > amount {
		return -1
	}
	if i := sort.SearchInts(coins, amount); i < len(coins) {
		if coins[i] == amount {
			return 1
		}

		coins = coins[:i]
	}

	dp := make([]int, amount+1)
	for _, c := range coins {
		dp[c] = 1
	}

	for i := coins[0] * 2; i <= amount; i++ {
		for _, c := range coins {
			if c > i {
				break
			}
			m, n := dp[c], dp[i-c]
			if m > 0 && n > 0 && (dp[i] == 0 || m+n < dp[i]) {
				dp[i] = m + n
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}
	return dp[amount]
}

func TestIt(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11

	if coinChange(coins, amount) != 3 {
		t.Fatal("bad answer")
	}
}
