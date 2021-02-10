package answer

import (
	"testing"
)

func maxTurbulenceSize(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return n
	}

	longest := 1
	start := 0
	nextShouldBigger := arr[0] > arr[1]

	for i := 0; i < n-1; i++ {
		if nextShouldBigger && arr[i] > arr[i+1] ||
			!nextShouldBigger && arr[i] < arr[i+1] {
			nextShouldBigger = !nextShouldBigger
		} else {
			length := i - start + 1
			if length > longest {
				longest = length
			}

			if arr[i] == arr[i+1] {
				for i < n-1 && arr[i] == arr[i+1] {
					i++
				}
				if i < n-1 {
					nextShouldBigger = arr[i] < arr[i+1]
				}
			}

			start = i
		}
	}

	if length := n - start; length > longest {
		longest = length
	}

	return longest
}

func TestMaxTurbulenceSize(t *testing.T) {
	testcases := [][]int{
		{9, 4, 2, 10, 7, 8, 8, 1, 9},
		{9, 4, 2, 10, 7, 8},
		{4, 8, 12, 16},
		{100},
		{9, 9},
		{0, 1, 1, 0, 1, 0, 1, 1, 0, 0},
		{37, 199, 60, 296, 257, 248, 115, 31, 273, 176},
	}
	vals := []int{
		5,
		5,
		2,
		1,
		1,
		5,
		5,
	}

	for i := 0; i < len(testcases); i++ {
		testcase := testcases[i]
		val := vals[i]
		res := maxTurbulenceSize(testcase)
		if val != res {
			t.Errorf("%v %v", testcase, i)
			t.Errorf("%v: return %v, but expect %v", testcase, res, val)
		}
	}
}

func TestMaxTurbulenceSize2(t *testing.T) {
	testcases := [][]int{
		{9, 9},
	}
	vals := []int{
		1,
	}

	for i := 0; i < len(testcases); i++ {
		testcase := testcases[i]
		val := vals[i]
		res := maxTurbulenceSize(testcase)
		if val != res {
			t.Errorf("%v %v", testcase, i)
			t.Errorf("%v: return %v, but expect %v", testcase, res, val)
		}
	}
}
