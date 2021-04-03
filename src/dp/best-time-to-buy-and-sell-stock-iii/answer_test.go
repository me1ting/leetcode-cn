package answer

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
// 依然是使用贪婪法+动态规划，核心是定义状态和状态转移方程
//
// 贪婪法是因为每一天的状态仅仅取决于前一天的状态，而且按照局部最优进行操作，基于股票的低买高卖的性质
//
// 每一天有以下几个状态：
// 未进行任何操作，收益为0
// 进行了一次买
// 进行了一次完整交易
// 进行了一次完整交易，然后进行了一次买
// 进行了两次完整交易
//
// 初始状态：
// 进行一次买的收益为-prices[0]，
// 进行一次完整交易的收益初始化为0（有效的完整收益必然>0，因此这里初始化为0没有任何问题）
// 进行一次完整交易，然后买的收益为-prices[0]，
//    这里比较绕，假设第一次买的价格为a, 第一次卖的价格为b，第二次买的价格为c，那么b>c，因为高卖低买
//    那么收益-a+b-c>-a是成立的，所以这里初始化为-prices[0]没有任何问题
// 进行两次完整交易的收益为0（有效的完整收益必然>0，因此这里初始化为0没有任何问题）
//
// （可以扩展为求最大允许N次交易的最大收益）
// 由于正确的初始化了第2,3..N次交易的收益，那么最终只需要返回第N次交易的最大收益
func maxProfit(prices []int) int {
	buys := []int{-prices[0], -prices[0]}
	sells := []int{0, 0}
	for i := 1; i < len(prices); i++ {
		buys[0] = max(buys[0], -prices[i])
		sells[0] = max(sells[0], buys[0]+prices[i])
		buys[1] = max(buys[1], sells[0]-prices[i])
		sells[1] = max(sells[1], buys[1]+prices[i])
	}
	return sells[1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
