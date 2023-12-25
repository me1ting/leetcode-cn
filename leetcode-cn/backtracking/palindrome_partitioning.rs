/// https://leetcode-cn.com/problems/palindrome-partitioning/
/// 参考题解：https://leetcode-cn.com/problems/palindrome-partitioning/solution/fen-ge-hui-wen-chuan-by-leetcode-solutio-6jkv/
/// 考察技巧：回溯、动态规划、记忆化，本题是考察这些技巧的难度适宜的题型
/// 回溯：
///    当我们考虑位置i为起点时，假设位置i之前的都是已经确定的回文字符串，我们会从i,i+1,i+2,..n-1作为j尝试每一种可能的回文字符串情况
///    并将其作为基础，递归调用j+1为起点的分割方法，每个递归调用的终点就是将当前解决方案复制保存到最终答案中去。
///    递归返回后，我们需要移除(i,j)的回文字符串记录，这样我们就可以共用同一个记录列表，来找到所有的解决方案。
///
/// 动态规划：通过构建表来避免重复计算。
///    我们在判断(i,j)是否是回文字符串时依赖于：palindrome(i,j) = s[i] == s[j] && palindrome(i+1,j-1);
///    而palindrome就可以使用动态规划
/// 记忆化：将需要重复计算的结果存储起来，需要时直接查表使用。与动态规划不同，记忆化只记录需要计算的中间值，而不是预先计算所有中间值。
///    使用记忆化可以对动态规划进行优化。
///
/// 从本题来讲，记忆化的逻辑判断可能不一定比动态规划在测试用例上表现更好。

struct Solution();

impl Solution {
    pub fn partition(s: String) -> Vec<Vec<String>> {
        let s = s.into_bytes();
        let n = s.len();

        let mut dp = vec![vec![true; n]; n];
        for j in 1..n {
            for i in 0..=(j - 1) {
                dp[i][j] = s[i] == s[j] && dp[i + 1][j - 1];
            }
        }

        let mut parts: Vec<(usize, usize)> = vec![];
        let mut ans: Vec<Vec<(usize, usize)>> = vec![];

        let mut sol = SolutionInner { dp, parts, ans, n };

        sol.do_part(0);

        sol.ans
            .iter()
            .map(|parts| {
                parts
                    .iter()
                    .map(|r| {
                        let (i, j) = *r;
                        String::from_utf8(s[i..=j].to_vec()).unwrap()
                    })
                    .collect()
            })
            .collect()
    }
}

struct SolutionInner {
    dp: Vec<Vec<bool>>,
    parts: Vec<(usize, usize)>,
    ans: Vec<Vec<(usize, usize)>>,
    n: usize,
}

impl SolutionInner {
    fn do_part(&mut self, i: usize) {
        if i == self.n {
            self.ans.push(self.parts.clone());
            return;
        }

        for j in i..self.n {
            if self.dp[i][j] {
                self.parts.push((i, j));
                self.do_part(j + 1);
                self.parts.pop();
            }
        }
    }
}
