package main
// https://leetcode-cn.com/problems/height-checker/
// 题目限定了数组元素和长度为1~100，因此使用桶排序的记录方式即可
func heightChecker(heights []int) int {
	count:=make([]int,101)
	for _,h:=range heights {
		count[h]++
	}

	moveTime:=0
	for i,k:=1,0;k<len(heights);i++ {
		for j:= count[i];j>0;j-- {
			if heights[k] != i {
				moveTime++
			}
			k++
		}
	}

	return moveTime
}
