package numberofwaystosplitarray

import "testing"

//https://leetcode.cn/problems/number-of-ways-to-split-array
func waysToSplitArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	res := 0
	leftSum := 0
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		if leftSum<<1 >= sum {
			res++
		}
	}
	return res
}

func TestWaysToSplitArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 4, -8, 7}, 2},
		{[]int{2, 3, 1, 0}, 2},
		{[]int{0, 0, 0, 0, 0}, 4},
	}

	for _, test := range tests {
		if result := waysToSplitArray(test.nums); result != test.expected {
			t.Errorf("For nums %v, expected %d, but got %d", test.nums, test.expected, result)
		}
	}
}
