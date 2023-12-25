pub struct Solution {}

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 这题既有动态规划，又有贪婪法
// 贪婪法：每一天的最大利润取决于前一天，也就是前一天的局部最优等于全局最优
// 动态规划：每天的最大利润是一种状态，通过一天一天的状态转移得到
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        use std::cmp;

        let mut profit = vec![0, 0]; //持有股票的最大利润，未持有股票的最大利润
        profit[0] = -prices[0];
        for i in 1..prices.len() {
            let with_max = cmp::max(profit[0], profit[1] - prices[i]);
            let without_max = cmp::max(profit[1], profit[0] + prices[i]);
            profit[0] = with_max;
            profit[1] = without_max;
        }
        return profit[1];
    }
}

#[cfg(test)]
mod tests {
    use crate::greedy::best_time_to_buy_and_sell_stock_ii;
    #[test]
    fn it_works() {
        let prices = vec![1, 2, 3, 4, 5];
        let expected = 4;
        assert_eq!(
            expected,
            best_time_to_buy_and_sell_stock_ii::Solution::max_profit(prices)
        );

        let prices = vec![7, 6, 4, 3, 1];
        let expected = 0;
        assert_eq!(
            expected,
            best_time_to_buy_and_sell_stock_ii::Solution::max_profit(prices)
        );
    }
}
