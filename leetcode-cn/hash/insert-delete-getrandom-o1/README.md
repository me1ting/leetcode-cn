# 简介

[原题](https://leetcode.cn/problems/insert-delete-getrandom-o1/)

实现一个`RandomizedSet`，要求查找和删除的时间复杂度为O(1)，并提供随机的O(1)的访问功能。

# 解析

最开始拿到该题，我以为是从头实现一个hashmap/hashset，埋头写了一大堆代码，后来才发现没法实现时间复杂度为O(1)的随机访问功能。

参考[题解](https://leetcode.cn/problems/insert-delete-getrandom-o1/solution/o1-shi-jian-cha-ru-shan-chu-he-huo-qu-su-rlz2/)要实现随机访问功能，需要使用数组来实现。

使用一个数组来记录数据：

- 添加值，添加到末尾
- 删除值，使用末尾元素替代被删除的值，数组长度-1

使用语言标准库提供的hashmap维护值与值在数组中的索引的映射。