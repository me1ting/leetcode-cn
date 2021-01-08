package answer

import "testing"

// https://leetcode-cn.com/problems/rotate-array/submissions/
// 最经典的自然是反转法
// 尝试模拟法来实现，浪费了大量时间和脑细胞也没能成功
func rotate(nums []int, k int) {
	k = k % len(nums)

	if k == 0 {
		return
	}

	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func TestIt(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, 3)

	for i, n := range nums {
		if n != expected[i] {
			t.Errorf("expected %v, return %v", expected, nums)
			break
		}
	}

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	expected = []int{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	rotate(nums, 38)

	for i, n := range nums {
		if n != expected[i] {
			t.Errorf("expected %v, return %v", expected, nums)
			break
		}
	}
}
