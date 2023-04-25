# 简介
题目：https://leetcode.cn/problems/flip-string-to-monotone-increasing
官方题解：https://leetcode.cn/problems/flip-string-to-monotone-increasing/solution/jiang-zi-fu-chuan-fan-zhuan-dao-dan-diao-stjd/

# 最少翻转使得字符串单调递增
存在元素只为0,1的序列，如何翻转位，使得序列是递增的。

使用动态规划，我们使用`dp[i][0]`,`dp[i][1]`分别记录使得`num[0:i+1]`在为0时所需要的最小翻转次数使得序列递增，为1时所需要的最小翻转次数使得序列递增：

```go
dp[i][0] = dp[i-1][0] //之前必然全为0
if num[i]==1 {
    dp[i][0]++ //需要翻转才能变成0
}

dp[i][1] = min(dp[i-1][0], dp[i-1][1]) //之前可能为0，或者为1 
if num[i]==0 {
    dp[i][0]++ //需要翻转才能变成1
}
```
其初始条件为：
```go
if num[i] == 1 {
    dp[i][0] = 1
}else {
    dp[i][1] = 1
}
```