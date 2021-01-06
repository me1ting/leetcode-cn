package squares_of_a_sorted_array

import "testing"

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
// 归并排序合并阶段的变种
func sortedSquares(A []int) []int {
	out := make([]int, len(A))
	pos := 0

	l, r := 0, len(A)-1
	c := r / 2

	if A[l] >= 0 {
		l = -1
		r = 0
	} else if A[r] <= 0 {
		l = r
		r = len(A)
	} else {
		for l+1 < r {
			c = l + (r-l)/2
			if A[c] < 0 {
				l = c
			} else if A[c] > 0 {
				r = c
			} else {
				break
			}
		}
		if A[c] == 0 {
			l = c - 1
			r = c
		}
	}

	for l >= 0 && r < len(A) {
		if -A[l] < A[r] {
			out[pos] = A[l] * A[l]
			l--
		} else {
			out[pos] = A[r] * A[r]
			r++
		}
		pos++
	}

	for l >= 0 {
		out[pos] = A[l] * A[l]
		l--
		pos++
	}
	for r < len(A) {
		out[pos] = A[r] * A[r]
		r++
		pos++
	}
	return out
}

func TestSquares(t *testing.T) {
	testCases := [][]int{
		[]int{-4, -1, 0, 3, 10},
		[]int{0, 2},
		[]int{-4, -3, -2, 3, 3},
	}
	expecteds := [][]int{
		[]int{0, 1, 9, 16, 100},
		[]int{0, 4},
		[]int{4, 9, 9, 9, 16},
	}
out:
	for j := 0; j < len(testCases); j++ {
		testCase := testCases[j]
		expected := expecteds[j]
		result := sortedSquares(testCase)
		for i := 0; i < len(testCase); i++ {
			if result[i] != expected[i] {
				t.Errorf("expected %v, recived %v", expected, result)
				break out
			}
		}
	}
}
