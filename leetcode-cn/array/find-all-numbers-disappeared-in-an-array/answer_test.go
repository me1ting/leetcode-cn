package answer

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/
// 要想不使用额外空间，那么使用原数组即可，使用负数表示统计数据
// 官方题解使用nums[index]/n表示统计个数，使用nums[index]%n表示其原来的值，更加简洁
func findDisappearedNumbers(nums []int) []int {
	indexOf := func(i int) int {
		return i - 1
	}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n <= 0 {
			continue
		}

		for nums[i] > 0 {
			n := nums[i]
			index := indexOf(n)
			if nums[index] > 0 {
				// 不同时赋值是避免index == i的情况
				nums[i] = nums[index]
				nums[index] = 0
			} else {
				nums[i] = 0
			}

			nums[index]--
		}
	}

	ans := []int{}
	for i, n := range nums {
		if n == 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}

func TestFindDisappearedNumbers(t *testing.T) {
	numss := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
		{1},
	}
	expecteds := [][]int{
		{5, 6},
		{},
	}

	for i := range numss {
		nums := numss[i]
		expected := expecteds[i]
		val := findDisappearedNumbers(nums)

		sort.Ints(expected)
		sort.Ints(val)

		if len(val) != len(expected) {
			t.Errorf("return %v, expected %v", val, expected)
		}

		for j := range val {
			if val[j] != expected[j] {
				t.Errorf("return %v, expected %v", val, expected)
			}
		}
	}
}
