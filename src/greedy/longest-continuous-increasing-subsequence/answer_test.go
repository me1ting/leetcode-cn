package answer

import "testing"

// https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/
// 一次遍历法？
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	start := 0
	max := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			length := i - start
			if length > max {
				max = length
			}
			start = i
		}
	}

	length := len(nums) - start
	if length > max {
		max = length
	}

	return max
}

func TestFindLengthOfLCIS(t *testing.T) {
	testcase := []int{2, 2, 2, 2}
	expected := 1
	val := findLengthOfLCIS(testcase)
	if val != expected {
		t.Errorf("expected %v, return %v", expected, val)
		return
	}

}
