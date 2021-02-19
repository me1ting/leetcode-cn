package max_consecutive_ones_iii

import "testing"

// https://leetcode-cn.com/problems/max-consecutive-ones-iii/
// 滑动窗口
// 基本思想：使用左右指针，指向全1（包括翻转）子序列的左右边界[)
// 右指针在右向移动时，遇到几种情况：
// - 为1，继续移动
// - 为0，考虑是否能够翻转，如果能翻转则继续移动，否则此时为当前子序列的最大长度，比较并记录，然后移动左指针，
// 直到释放一个翻转机会，继续移动右子针
//
// 优化：
// 对于最大值情况，我们初始得到一个满足条件的最大窗口，在后续情况时会遇到不符合条件的情况
// 我们不需要缩小窗口以满足条件，只需要保持窗口大小即可，即在移动右指针的同时移动左指针
func longestOnes(A []int, K int) int {
	l, r := 0, 0
	holeCount := 0

	for _,n := range A {
		r++
		if n==0 {
			holeCount++
		}
		if holeCount>K {
			if A[l] == 0 {
				holeCount--
			}
			l++
		}
	}

	return r-l
}

func TestLongestOnes(t *testing.T) {
	As := [][]int{
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	Ks := []int{
		2, 3, 0, 0, 0,
	}
	expecteds := []int{
		6, 10, 4, 0, 5,
	}

	for i, A := range As {
		K := Ks[i]
		expected := expecteds[i]

		ans := longestOnes(A, K)
		if ans != expected {
			t.Errorf("A=%v,K=%v, return %v, expected %v", A, K, ans, expected)
			break
		}
	}
}
