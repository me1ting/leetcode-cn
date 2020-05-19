package main

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
// 排序数组，缺失唯一整数，使用二分搜索即可
// 从实现上来讲还可以更加优雅，可以查看示例
func missingNumber(nums []int) int {
	if nums[0] > 0 {
		return 0
	}

	if nums[len(nums)-1] == len(nums)-1{
		return len(nums)
	}

	l,r:=0,len(nums)-1 // num[l] == l,nums[r] > r
	for l<r-1 {
		c:=(l+r)/2
		if nums[c] == c {
			l = c
		}else if nums[c] > c {
			r = c
		}
	}

	return r
}

func main() {
	missingNumber([]int{0,2,3})
}