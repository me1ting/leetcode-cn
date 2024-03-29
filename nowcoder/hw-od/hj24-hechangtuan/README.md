# 合唱团

原题[地址](https://www.nowcoder.com/practice/6d9d69e3898f45169a441632b325c7b4)

N 位同学站成一排，音乐老师要请最少的同学出列，使得剩下的 K 位同学排成合唱队形。

通俗来说，能找到一个同学，他的两边的同学身高都依次严格降低的队形就是合唱队形。
例子：
123 124 125 123 121 是一个合唱队形
123 123 124 122不是合唱队形，因为前两名同学身高相等，不符合要求
123 122 121 122不是合唱队形，因为找不到一个同学，他的两侧同学身高递减。

你的任务是，已知所有N位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。

注意：不允许改变队列元素的先后顺序 且 不要求最高同学左右人数必须相等

## 解析

可以用dp来解决问题，使用两次dp：

- `dp1[i]`表示以i为起点的最长递增子序列
- `dp2[i]`表示以i为终点的最长递增子序列，从右向左构建

最后遍历数组，统计`Max(dp1[i]+dp2[i])`，即为结果。