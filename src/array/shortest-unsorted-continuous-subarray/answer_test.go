package shortestunsortedcontinuoussubarray

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	numss := [][]int{
		{2, 6, 4, 8, 10, 9, 15},
		{1, 2, 3, 4},
		{1},
		{1, 5, 6, 7, 8, 9, 2, 3, 4},
		{1, 2, 3, 7, 4, 5, 6, 8},
	}
	expecteds := []int{
		5,
		0,
		0,
		8,
		4,
	}

	for i, nums := range numss {
		expected := expecteds[i]
		val := findUnsortedSubarray(nums)
		if val != expected {
			t.Errorf("expected %v, return %v", expected, val)
		}
	}
}
