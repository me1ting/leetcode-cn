package answer

import "sort"

// https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
// 全是非负数/非正数，最大的数相乘
// 存在负数，两个最小负数*最大正数

// 标准题解实现很巧妙，不需要判定过多的条件，直接计算两个可能值，取最大值
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
