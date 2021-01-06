package patchingarray

import "testing"

// https://leetcode-cn.com/problems/patching-array/
// 贪婪法
func minPatches(nums []int, n int) int {
	// ToDo
	return 0
}

func TestIt(t *testing.T) {
	input1 := [][]int{
		{1, 3},
		{1, 5, 10},
		{1, 2, 2},
	}
	input2 := []int{
		6,
		20,
		5,
	}

	expecteds := []int{
		1,
		2,
		0,
	}

	for i := 0; i < len(input1); i++ {
		if minPatches(input1[i], input2[i]) != expecteds[i] {
			t.Errorf("error")
		}
	}
}
