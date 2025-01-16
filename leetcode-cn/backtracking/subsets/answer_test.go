package subsets

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/subsets/solutions/420294/zi-ji-by-leetcode-solution/
// 有两种思路来解决这个问题：
//
// 思路1：将一个数字想象成0,1，0表示解不包含该数字，1表示解包含该数字，那么解的集合，可以映射为[0,2^len(nums))。
// 因此可以遍历[0,2^len(nums))，并将每个数字转换为一个解。
//
// 思路2：首先我们访问nums[0]，有两种情况：包含nums[0]，不包含nums[0]，分别在两种情况下，再访问nums[1]...
// 直到我们访问完所有的数字，我们使用一个Set()来记录包含的元素，并在到达末尾时得到一个解。
func subsets(nums []int) [][]int {
	res := [][]int{}
	set := make([]int, 0, len(nums))

	var visit func(i int)
	visit = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, set...))
			return
		}
		set = append(set, nums[i])
		visit(i + 1)
		set = set[:len(set)-1]
		visit(i + 1)
	}

	visit(0)
	return res
}

func TestSubsets(t *testing.T) {
	log.Println(subsets([]int{1, 2, 3}))
}
