package main

import "sort"

func longestStrChain(words []string) int {
	N := len(words)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	dp := make([]int, N)
	dp[0] = 1
	for i := 1; i < N; i++ {
		dp[i] = 1
		preLen := len(words[i]) - 1
		for j := i - 1; j >= 0 && len(words[j]) >= preLen; j-- {
			if len(words[j]) == preLen && isPre(words[j], words[i]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLen := 1
	for i := 0; i < N; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

// test a is b's pre.
func isPre(a, b string) bool {
	if len(a)+1 != len(b) {
		return false
	}
	i, j := 0, 0

	for i < len(a) && j < len(b) && a[i] == b[j] {
		i++
		j++
	}
	j++
	for i < len(a) && j < len(b) && a[i] == b[j] {
		i++
		j++
	}
	return i == len(a) && j == len(b)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
