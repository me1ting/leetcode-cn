# 简介

https://leetcode.cn/problems/3sum/

# 解析

设满足条件的三个值其索引位置从左往右为：i,j,k；其值为x,y,z

## 二分查找

一开始我的思路是从左网右遍历每个可能的i, 再遍历每个可能的j，使用二分查找值为z的k。

提交代码后发现运行效率很慢。

## 基于双指针的two-sum

最佳的方式是对于每一个可能的i，在`nums[i+1:]`查找满足`nums[j]+nums[k]==-nums[i]`的j,k，转换为two-sum问题。

在不引入额外存储空间的情况下，two-sum的最佳解法是使用双指针法（时间复杂度O(n)）：

```go
// 为了解释双指针法，这里假设nums是一个不包含重复元素的数组
for j,k:=0,n;j<k; {
    sum := nums[j]+nums[k]
    if sum < 0 {
        j++
    }else if sum > 0 {
        k--
    }else{
        // 找到了，并记录
        j++
        k--
    }
}
```