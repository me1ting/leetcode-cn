struct Solution {}

// https://leetcode-cn.com/problems/house-robber-ii/
// 直接想法：max(sum_odd,sum_even)，但会遇到特殊情况如：0 1 0 1 0 1 9 1，最大值是1+1+9并不同时位于奇数项或者偶数项上
// 再看此题，很像动态规划：$max_i = MAX(max_{i-2}+nums[i],max_{i-1})$
// 由于环绕性质，还要考虑边界情况，解决办法是分别求解nums[1..],nums[..len-1]的然后取最大值
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        use std::cmp::max;
        if nums.len() == 1 {
            return nums[0];
        }
        if nums.len() == 2 {
            return max(nums[0], nums[1]);
        }

        let mut dp = vec![0; nums.len()];
        dp[0] = nums[0];
        dp[1] = max(nums[0], nums[1]);

        for i in 2..nums.len() - 1 {
            dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);
        }

        let without_last = dp[nums.len() - 2];

        dp[1] = nums[1];
        dp[2] = max(nums[1], nums[2]);

        for i in 3..nums.len() {
            dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);
        }
        let without_first = dp[nums.len() - 1];

        max(without_first, without_last)
    }
}
