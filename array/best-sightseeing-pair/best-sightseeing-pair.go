package main
// https://leetcode-cn.com/problems/best-sightseeing-pair/
// 目标是O(n)，通过记录评分衰减即可实现
func maxScoreSightseeingPair(A []int) int {
	sum,max:=0,1 //当前最大评分组合，当前最具有吸引力的景点
	for _,v:= range A {
		max--
		if max + v > sum {
			sum = max + v
		}
		if v > max {
			max = v
		}
	}
	return sum
}
