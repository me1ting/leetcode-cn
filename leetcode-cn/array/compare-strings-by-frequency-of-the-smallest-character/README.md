# 简介

[原题](https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/)

# 解析

这题描述的有点复杂，可以简化为：

- 函数`f()`计算字符串中，最小字符串的出现次数
- 对于`queries`中的每一个字符串，计算在`words`中，其`f(word)`值大于`f(query)`的word的个数

## 后缀和数组

`后缀和数组`术语用于描述一个避免重复计算的数组，对于数组`a`，其后缀和数组`sum[i] = a[i]+a[i+1]+...+a[n-1]`。

相应的，也有`前缀和数组`概念。