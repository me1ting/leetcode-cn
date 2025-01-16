package shortestsubarraywithoratleastki

import (
	"math"
	"testing"
)

// https://leetcode.cn/problems/shortest-subarray-with-or-at-least-k-i/description/
// 使用滑动窗口，维护一个val，表示窗口内的所有元素的或值，维护一个bits，表示窗口内每个元素的二进制位1的个数
// 当val<k时，右指针右移，否则左指针右移，直到val<k，记录最小长度
func minimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}

	length := len(nums)

	res := math.MaxInt32
	bits := make([]int, 7) //0<=nums[i]<=50
	val := 0

	for l, r := 0, 0; r < length; r++ {
		n := nums[r]
		val |= n
		for i := range 7 {
			if n&(1<<i) > 0 {
				bits[i]++
			}
		}
		if val < k {
			continue
		}

		for val >= k && l < r {
			n := nums[l]
			for i := range 7 {
				if n&(1<<i) > 0 {
					bits[i]--
					if bits[i] == 0 {
						val &= ^(1 << i)
					}
				}
			}
			l++
		}
		if val >= k {
			return 1
		} else {
			res = min(res, r-l+2)
		}
	}

	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func TestMinimumSubarrayLength(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3}, 2, 1},
		{[]int{2, 1, 8}, 10, 3},
		{[]int{1, 2}, 0, 1},
		{[]int{16, 1, 2, 20, 32}, 45, 2},
	}

	for _, test := range tests {
		result := minimumSubarrayLength(test.nums, test.k)
		if result != test.expected {
			t.Errorf("For nums %v and k %d, expected %d but got %d", test.nums, test.k, test.expected, result)
		}
	}
}
