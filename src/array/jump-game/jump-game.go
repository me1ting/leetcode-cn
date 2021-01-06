package main

// https://leetcode-cn.com/problems/jump-game/
/*
// 这是我第一时间想到的方法，但个人思考遇到极端情况会导致该算法极端性能下降
// "回退避坑法"：
// 尽量跳最远距离，如果当前位置超过last，则说明可以跳到last
// 如果遇到坑，从坑开始倒退检查，能否越过坑，不能越过则返回false；能越过则更新位置
func canJump(nums []int) bool {
	pos ,last:= 0,len(nums) - 1
	out:
	for pos < last {
		step := nums[pos]
		if nums[pos] == 0 {
			for try:=pos-1; try>=0; try-- {
				if nums[try] > pos - try{
					pos = try + nums[try]
					continue out
				}
			}
			return false
		}else{
			pos += step
		}
	}
	return true
}
*/

// BFS算法
// 从实际效果来看BFS的执行时间过长
func canJump(nums []int) bool {
	queue := make([]int,len(nums))
	queue[0] = 0
	nums[0] = - nums[0] //visited
	l,r := 0,1 //first element pos, to be written pos
	for l < r {
		n := queue[l]
		step := -nums[n]
		if n+step >= len(nums) - 1{
			return true
		}
		for i:= n+1; i<=n+step ; i++{
			if nums[i] > 0 {
				nums[i] = -nums[i] //visited
				queue[r] = i
				r++
			}
		}
		l++
	}
	return false
}

/*
// O(n)也是最简洁、容易理解的实现。单次扫描，记录最远距离，如果i不在覆盖范围，说明无法到达。如果i可以更新最远距离，则更新。
// 如果最远距离等于或超过最终位置，说明成功。
func canJump(nums []int) bool {
	rightMost := 0
	last := len(nums) -1
	for i,n := range nums {
		if i > rightMost {
			return false
		}

		if most:=i+n;most>rightMost{
			rightMost = most

			if rightMost >=  last{
				return true
			}
		}
	}
	return true
}
*/

func main() {
	canJump([]int{2, 5, 0, 0})
}
