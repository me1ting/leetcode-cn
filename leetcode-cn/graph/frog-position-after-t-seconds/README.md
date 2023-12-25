# 简介
原题：https://leetcode.cn/problems/frog-position-after-t-seconds/submissions/

# 解析
使用搜索算法来实现，dfs更简单一点。

answer_test.go的代码还可以进一步优化：

- 可以将时间t传递到搜索过程，当深度大于t时，返回父节点而不继续搜索更深的节点
- 找到节点后，可以直接返回结果，不需要完整的搜索完图