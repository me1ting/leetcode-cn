package best_time_to_buy_and_sell_stock_with_transaction_fee

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-han-sh-rzlz/
// 动态规划法：
// 重点是理解如何将本题转换为动态规划解法：
// 1.定义状态
// 2.确定转移方程
// 3.迭代计算出状态表
// 4.返回状态表中的最后结果
//
// 1.定义状态：每天有两个状态：未持有股票时的最大利润，持有股票时的最大利润
// 2.确定转移方程：
//    当天持有股票的最大利润，可能是继承自昨天持有股票时的最大利润，或者昨天未持有股票，今天买入
//    当天未持有股票的最大利润，可能是继承自昨天未持有股票时的最大利润，或者昨天持有股票，今天卖出
// 注意的是上述状态、状态转移方程可能与我们在思考这个问题时的直观思维有很大出入，这需要适应。
// 实质上来讲，这如同于读书时候做的数学计算题，需要死记掌握解法

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
