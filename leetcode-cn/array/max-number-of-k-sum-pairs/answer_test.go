package maxnumberofksumpairs

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	count := 0
	for i, j := 0, n-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == k {
			count++
			i++
			j--
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return count
}
