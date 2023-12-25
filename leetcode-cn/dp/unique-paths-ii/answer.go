package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m+1)
	//为了使得dp[1][1] = 1
	dp[1] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[j] = dp[j-1] + dp[j]
			if obstacleGrid[j-1][i-1] == 1 {
				dp[j] = 0
			}
		}
	}

	return dp[m]
}
