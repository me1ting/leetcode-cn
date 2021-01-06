package main

// https://leetcode-cn.com/problems/find-the-duplicate-number
// https://leetcode-cn.com/problems/find-the-duplicate-number/solution/xun-zhao-zhong-fu-shu-by-leetcode-solution/

// 二分查找
// 只有一个重复数，且余１，则小于等于该数的count必然大于该数
func findDuplicate(nums []int) int {
	left,right:=0, len(nums)

	for left < right {//count[left]<=left,count[right]>right，即left一定不是，right可能是
		count := 0
		center := (left+right)/2
		for _,n:= range nums {
			if n<=center {
				count++
			}
		}
		if count <= center {//一定不是
			left = center + 1
		}else{//可能是
			right = center
		}
	}
	return left
}