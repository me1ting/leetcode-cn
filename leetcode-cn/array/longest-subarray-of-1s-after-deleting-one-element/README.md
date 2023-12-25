# 简介

原题：https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/

# 解析

```go
func longestSubarray(nums []int) int {
	ans := 0
	len1, len2 := 0, 0
	for _, num := range nums {
		if num == 0 {
			ans = max(ans, len1+len2)
			len1 = len2
			len2 = 0
		} else {
			len2++
		}
	}
	ans = max(ans, len1+len2)

	if ans == len(nums) {
		ans--
	}
	return ans
}
```

我们可以使用len1和len2记录被0分割的均为1的子字符串长度，我们遍历每一个字符：

- 当遇到0时，我们会得到一个结果(len1+len2)，并更新结果。然后更新len1为len2，重置len2为0。
- 当遇到1时，len2++

边界情况：

- 字符串可能以1结尾，我们需要将其统计为结果，并更新结果。
- 字符串可能全为1，我们需要至少删除一个字符。
