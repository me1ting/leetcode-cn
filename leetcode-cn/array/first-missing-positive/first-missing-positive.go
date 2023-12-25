package main

// https://leetcode-cn.com/problems/first-missing-positive/
// 可以利用数组的长度：
// 最小整数最大为len+1
// n>len,n<=0，丢弃，否则存储在n-1的位置
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	for i, l := 0, len(nums); i < l; {
		n := nums[i]
		if n > l || n < 1 {
			nums[i] = 0//使用0标记为不存在
			i++
		} else if n == i+1 {//位置i存i+1的值，从而表示索引
			i++
		} else {
			if nums[n-1] == n {// 避免swap死循环
				nums[i] = 0
			}else{
				nums[i], nums[n-1] = nums[n-1], nums[i]
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return len(nums) + 1 // 极端情况，最大值为len+1
}
