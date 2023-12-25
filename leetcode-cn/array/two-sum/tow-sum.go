package main

// https://leetcode-cn.com/problems/two-sum/
// 没有说明元素是否排序，是否只为正数，那就是未排序，存在正负数
// 使用map解决问题即可
func twoSum(nums []int, target int) []int {
	exist := make(map[int]int)
	for i,n := range nums {
		if index,ok := exist[target - n];ok{
			return []int{index,i}
		}
		exist[n] = i
	}
	return nil
}
