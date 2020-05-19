package main

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/submissions/
// 简单问题，取决于原地检查还是使用map
// 原地检查最多迭代Ｎ次，交换Ｎ次，因此是Ｏ(n)时间复杂度。
func findRepeatNumber(nums []int) int {
	for i:=0;i< len(nums);i++{
		n := nums[i]
		for n != i {
			if nums[n] == n {
				return n
			}
			nums[i],nums[n] = nums[n],n
			n = nums[i]
		}
	}
	return -1
}

func main() {
	findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3})
}