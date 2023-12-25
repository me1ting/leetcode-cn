# 简介
原题：https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays

# 解析
这道题你可以用DP的思想来解释算法，但是最简单的理解是：

![](.assets/werwv342d.png)

对每一个可能的右子数组，我们只需要记录左子数组的最大值，结果为和的最大值。

由于本题没有限制先后顺序，我们需要交换左右子数组的长度，总共计算两次，取最大值。