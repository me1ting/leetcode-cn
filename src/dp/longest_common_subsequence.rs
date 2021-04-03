struct Solution();

// https://leetcode-cn.com/problems/longest-common-subsequence/
// 参考题解：https://leetcode-cn.com/problems/longest-common-subsequence/solution/zui-chang-gong-gong-zi-xu-lie-by-leetcod-y7u0/
//
// 需要使用动态规划，这里的dp[m][n]表示text1[:m],text2[:n]的最大公共字符串长度
impl Solution {
    pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
        let text1 = text1.into_bytes();
        let text2 = text2.into_bytes();
        let mut dp = vec![vec![0; text2.len() + 1]; text1.len() + 1];
        for (i, c1) in text1.iter().enumerate() {
            for (j, c2) in text2.iter().enumerate() {
                if c1 == c2 {
                    dp[i + 1][j + 1] = dp[i][j] + 1;
                } else {
                    dp[i + 1][j + 1] = std::cmp::max(dp[i + 1][j], dp[i][j + 1]);
                }
            }
        }
        dp[text1.len()][text2.len()]
    }
}
