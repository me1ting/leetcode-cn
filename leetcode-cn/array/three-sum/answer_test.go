package threesum

import (
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n-2; {
		x := nums[i]
		//two sum
		for j, k := i+1, n-1; j < k; {
			y, z := nums[j], nums[k]
			sum := y + z
			if sum < 0-x {
				j++
			} else if sum > 0-x {
				k--
			} else {
				ans = append(ans, []int{x, y, z})
				j++
				k--
				for j < k && nums[j] == y {
					j++
				}
				for j < k && nums[k] == z {
					k--
				}
			}
		}
		i++
		for i < n-2 && nums[i] == x {
			i++
		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	if len(threeSum(nums)) != 2 {
		t.Errorf("len of result should be 2")
	}
}
