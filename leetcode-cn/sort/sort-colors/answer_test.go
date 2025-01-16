package sortcolors

import (
	"reflect"
	"testing"
)

func sortColors(nums []int) {
	redCnt, whiteCnt, blueCnt := 0, 0, 0
	for _, color := range nums {
		if color == 0 {
			redCnt++
		} else if color == 1 {
			whiteCnt++
		} else {
			blueCnt++
		}
	}

	for i := 0; i < redCnt; i++ {
		nums[i] = 0
	}
	for i := redCnt; i < redCnt+whiteCnt; i++ {
		nums[i] = 1
	}
	for i := redCnt + whiteCnt; i < redCnt+whiteCnt+blueCnt; i++ {
		nums[i] = 2
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{2, 2, 2, 2, 2}, []int{2, 2, 2, 2, 2}},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}},
		{[]int{1, 0, 1, 0, 1, 0}, []int{0, 0, 0, 1, 1, 1}},
	}

	for _, test := range tests {
		sortColors(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, test.input)
		}
	}
}
