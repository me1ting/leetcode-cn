package answer

import (
	"fmt"
	"testing"
)

func summaryRanges(nums []int) []string {
	res := []string{}

	if len(nums) == 0 {
		return res
	}

	start := 0
	n := nums[0]
	i := 1

	for ; i < len(nums); i++ {
		n++
		if nums[i] != n {
			if i > start+1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			} else {
				res = append(res, fmt.Sprintf("%d", nums[start]))
			}
			start = i
			n = nums[i]
		}
	}

	if i > start+1 {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
	} else {
		res = append(res, fmt.Sprintf("%d", nums[start]))
	}

	return res
}

func TestIt(t *testing.T) {
	testcases := [][]int{
		{0, 1, 2, 4, 5, 7},
		{0, 2, 3, 4, 6, 8, 9},
		{},
		{0},
	}
	expecteds := [][]string{
		{"0->2", "4->5", "7"},
		{"0", "2->4", "6", "8->9"},
		{},
		{"0"},
	}

	for i, testcase := range testcases {
		expected := expecteds[i]
		res := summaryRanges(testcase)
		if len(res) != len(expected) {
			t.Errorf("expected %v, return %v", expected, res)
			break
		}
		for j, str := range res {
			if str != expected[j] {
				t.Errorf("expected %v, return %v", expected, res)
				break
			}
		}
	}
}
