package main

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
// 从朴素的思考上来讲，因为数组包含正数，只要遇到和为负数就丢弃，就可以解决问题
// 如果题目条件放开，允许全部负数，那么只需要额外记录最小负数，按照同样的方法可以解决问题

// 但是，这里为了学习动态规划，这里使用动态规划的思想来解决问题：
//
// - 状态定义 描述问题。在本题中只存在一维状态，使用dp[n]表示以nums[n]结尾的子数组的连续子数组最大合。
// - 转移方程 状态与子状态的关系。dp[i-1]<=0,dp[i]=nums[i];dp[i-1]>0,dp[i]=dp[i-1]+nums[i]
// - 初始状态 dp[0] = nums[0]
// - 返回值 dp中的最大值

func maxSubArray(nums []int) int {
	// 并不一定要创建dp数组
	sum := nums[0]
	max:=sum
	for i:=1;i< len(nums);i++{
		if sum <= 0 {
			sum = nums[i]
		}else{
			sum += nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
