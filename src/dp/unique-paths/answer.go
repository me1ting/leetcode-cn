package main

// 这里假定列为m，行为n
func uniquePaths(m int, n int) int {
	// 压缩后的dp，只记录前一行的记录
	dp := make([]int, m+1)
	// 第一列只有一条路径
	for j := 0; j <= m; j++ {
		dp[j] = 1
	}
	//从第二列，第二行开始推导
	for i := 2; i <= n; i++ {
		for j := 2; j <= m; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[m]
}
