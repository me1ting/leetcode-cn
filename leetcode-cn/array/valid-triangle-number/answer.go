package validtrianglenumber

import "sort"

// https://leetcode-cn.com/problems/valid-triangle-number/
// https://leetcode-cn.com/problems/valid-triangle-number/solution/you-xiao-san-jiao-xing-de-ge-shu-by-leet-t2td/
// 排序+双指针，见官方题解
func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	count := 0
	for i := 0; i < n; i++ {
		k := i + 2
		for j := i + 1; j < n; j++ {
			for k < n && nums[k] < nums[i]+nums[j] {
				k++
			}
			count += max(k-j-1, 0)
		}
	}
	return count
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
