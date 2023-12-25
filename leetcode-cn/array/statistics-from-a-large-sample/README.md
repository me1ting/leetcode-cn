# 简介
原题：https://leetcode.cn/problems/statistics-from-a-large-sample/

# 解析
这是一道编程题，考察：

- 细心，考虑边界情况
- 对基础数据类型的掌握，以Go为例，int根据平台可能为int32或者int64，数据的sum无法用int32表示，需要指定int64。整数类型的除法会取整，求平均值时要小心。