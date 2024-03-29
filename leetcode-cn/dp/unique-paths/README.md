# about
题目：https://leetcode.cn/problems/unique-paths/

# 分析
本题可以用数学在`O(1)`实际内回答，但是当题目对内容进行修改（见`unique-paths-ii`），比如某些路径无法通过时，这种取巧的办法就没用了。

可以使用动态规划解决问题，使用dp[i][j]表示到达第`i`列`j`行的不同路径数量，那么很容易想到：

- `dp[1][j] = 1`,`dp[i][1] = 1`
- `dp[i][j] = dp[i-1][j] + dp[i][j-1]`

这里假设我们一列一列的递推，那么数据只涉及前一列和当前列，因此可以压缩dp的空间。