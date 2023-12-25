package answer

import "testing"

// https://leetcode-cn.com/problems/binary-prefix-divisible-by-5/submissions/
func prefixesDivBy5(A []int) []bool {
	ans := []bool{}
	rem := 0
	for _, bit := range A {
		rem = (rem<<1 + bit) % 5
		res := false
		if rem == 0 {
			res = true
		}
		ans = append(ans, res)
	}
	return ans
}

func TestIt(t *testing.T) {
	testcases := [][]int{
		{0, 1, 1},
		{1, 1, 1},
		{0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 1},
	}
	expecteds := [][]bool{
		{true, false, false},
		{false, false, false},
		{true, false, false, false, true, false},
		{false, false, false, false, false},
	}

	for i, testcase := range testcases {
		expected := expecteds[i]
		result := prefixesDivBy5(testcase)
		for j := 0; j < len(expected); j++ {
			if result[i] != expected[i] {
				t.Errorf("expected %v,return %v", expected, result)
			}
		}
	}
}
