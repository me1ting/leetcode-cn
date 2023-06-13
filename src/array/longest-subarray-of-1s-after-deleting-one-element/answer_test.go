package longestsubarrayof1safterdeletingoneelement

import "testing"

func longestSubarray(nums []int) int {
	ans := 0
	len1, len2 := 0, 0
	for _, num := range nums {
		if num == 0 {
			ans = max(ans, len1+len2)
			len1 = len2
			len2 = 0
		} else {
			len2++
		}
	}
	ans = max(ans, len1+len2)

	if ans == len(nums) {
		ans--
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestLongestSubarray(t *testing.T) {
	numsList := [][]int{
		{0, 1, 1, 1, 0, 1, 1, 0, 1},
	}
	rightAnsList := []int{
		5,
	}

	for i, nums := range numsList {
		rightAns := rightAnsList[i]

		ans := longestSubarray(nums)

		if ans != rightAns {
			t.Errorf("expected %v,but got %v", rightAns, ans)
		}
	}
}
