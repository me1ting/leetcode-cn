package main

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 没啥难度，使用指针记录位置即可

func removeDuplicates(nums []int) int {
	pos:=1
	for i:=1;i<len(nums);i++{
		if nums[i] != nums[i-1] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
