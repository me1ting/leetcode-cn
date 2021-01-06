package main

// https://leetcode-cn.com/problems/jump-game-ii
/*
//BFS算法，性能不怎么样，但能解决问题
func jump(nums []int) int {
	if len(nums)==1 {
		return 0
	}
	queue := make([]int,len(nums))
	depth := make([]int,len(nums))

	queue[0] = 0
	depth[0] = 0

	nums[0] = - nums[0] //visited
	l,r := 0,1 //first element pos, to be written pos
	for l < r {
		n := queue[l]
		step := -nums[n]
		d := depth[n]
		if n+step >= len(nums) - 1{
			return d+1
		}
		for i:= n+1; i<=n+step ; i++{
			if nums[i] > 0 {
				depth[i] = d+1
				nums[i] = -nums[i] //visited
				queue[r] = i
				r++
			}
		}
		l++
	}
	return 0
}
*/

// 贪心算法
// 新覆盖返回时更新最小跳跃次数
// 步数技巧：不需要记录所有位置的最小跳跃步数，只需要记录两个边界：
// 当前步数范围，下一步数所能达到的最远距离（不断更新）
// 当i超过当前步数范围时，更新步数和其匹配的范围
func jump(nums []int) int {
	last := len(nums) -1
	currMost,lastMost,steps :=0,0,0
	for i:=0; i<last ;i++{
		if most:=i+nums[i] ; most>lastMost {
			lastMost = most
		}
		if i == currMost {
			currMost = lastMost
			steps++
		}
	}
	return steps
}

func min(x,y int) int {
	if x < y {
		return x
	}else{
		return y
	}
}
