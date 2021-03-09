struct Solution();

/// https://leetcode-cn.com/problems/palindrome-partitioning-ii/
/// 参考题解：https://leetcode-cn.com/problems/palindrome-partitioning-ii/solution/fen-ge-hui-wen-chuan-ii-by-leetcode-solu-norx/
/// 考察技巧：动态规划
///
/// 动态规划的题一般都比较难以理解，是因为状态方程很难通过直观的方式找到。
/// 对于0..i的子字符串，它有两种情况：
///
/// - 0..i是回文字符串，切0刀
/// - 0..i不是回文字符串，那么需要前缀的切次数+切末尾回文字符串的一刀，遍历找到最小的刀数
impl Solution {
    pub fn min_cut(s: String) -> i32 {
        let s = s.into_bytes();
        let n = s.len();

        let mut dp = vec![vec![true; n]; n];
        for j in 1..n {
            for i in 0..=(j - 1) {
                dp[i][j] = s[i] == s[j] && dp[i + 1][j - 1];
            }
        }

        let mut cuts = vec![0; n];

        for j in 0..n {
            if dp[0][j] {
                continue;
            }
            cuts[j] = j;
            for i in 1..=j {
                if dp[i][j] {
                    let curr = cuts[i - 1] + 1;
                    if curr < cuts[j] {
                        cuts[j] = curr;
                    }
                }
            }
        }

        cuts[n - 1] as i32
    }
}

#[cfg(test)]
mod tests {
    use crate::dp::palindrome_partitioning_ii::Solution;
    #[test]
    fn test_min_cut() {
        let ss = vec!["aaba", "a", "aab"];
        let expecteds = vec![1, 0, 1];

        for i in 0..ss.len() {
            let s = ss[i];
            let expected = expecteds[i];
            let ans = Solution::min_cut(String::from(s));
            assert_eq!(ans, expected);
        }
    }
}
