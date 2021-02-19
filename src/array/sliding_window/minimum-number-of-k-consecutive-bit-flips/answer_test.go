package answer

import "testing"

// https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/
// 参考官方题解 https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/solution/k-lian-xu-wei-de-zui-xiao-fan-zhuan-ci-s-bikk/
// 基本策略（贪婪法）：
//
// - 每个位置翻转两次不改变结果，因此每个位置只需要翻转一次，不需要特殊的技巧
// - 从左到右，看到0就进行翻转，如果最后0出现的位置没有足够的空间进行翻转，则无法达成目标，否则可以翻转得到全1，翻转次数即为最小次数
//
// 优化：**使用差分数组**
// 基本思想：我们可以使用额外的空间记录每个位置的翻转情况，但这等价于原地翻转
// 差分数组：使用一个变量记录当前位置的翻转次数，使用数组记录位置i与i-1翻转次数的差（rev(i)-rev(i-1)）
//			这样，每个位置的翻转次数 = 实际翻转次数 + 差分的累加，这个可以在遍历过程中完成
//			同时，每次翻转只影响diff[i]与diff[i+k]，这样达成了我们优化的目的：避免每次都需要进行k次读写操作
//
// 进一步优化：**滑动窗口**
// 基本思想：使用其它int值表示位置i开始的子数组发生翻转过，从而不需要分配额外的空间
//
// 使用位运算：
//          由于此题只考察翻转次数的奇偶性，因此可以使用位运算而不用记录实际的翻转次数
//
// 总结：
// 此题考察了两个思想：
// - 差分数组，优化数组读写
// - 如何原地承载状态
func minKBitFlips(A []int, K int) (ans int) {
	N := len(A)
	revCount := 0

	for i, n := range A {
		if i >= K && A[i-K] > 1 {
			revCount ^= 1
			//恢复原值，注释掉是因为对当前算法而言可以省略掉该操作
			//A[i-k] = 1
		}
		n ^= revCount

		if n == 0 {
			if i+K > N {
				return -1
			}
			ans++
			revCount ^= 1
			A[i] = 2
		}
	}

	return ans
}

func TestMinSwapsCouples(t *testing.T) {
	As := [][]int{
		{0, 1, 0},
		{1, 1, 0},
		{0, 0, 0, 1, 0, 1, 1, 0},
	}
	Ks := []int{
		1, 2, 3,
	}
	expecteds := []int{
		2, -1, 3,
	}

	for i, A := range As {
		K := Ks[i]
		expected := expecteds[i]
		ans := minKBitFlips(A, K)

		if ans != expected {
			t.Errorf("return %v, expected %v", ans, expected)
			break
		}
	}
}
