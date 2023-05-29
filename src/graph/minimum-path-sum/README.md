# 简介
原题：https://leetcode.cn/problems/minimum-path-sum/submissions/

# 解析

## 方法1：单源有权最短路径

拿到这一题时，我的第一思路是使用dijkstra算法来求单源有权最短路径。但是实现复杂，效率不高。

## 方法2：动态规划

参考题解，可以使用动态规划求解。根据题目可以得知，只能向右或者向下走，得到转移方程：

`paths[row][col] = min(paths[row-1][col],paths[row][col-1])+grid[row][col]`

