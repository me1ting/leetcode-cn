package answer

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/sort-integers-by-the-power-value/solution/jiang-zheng-shu-an-quan-zhong-pai-xu-by-leetcode-s/
// 考察使用记忆化的技巧
//
// 可选优化：可以使用快速选择算法来不加快计算速度
func getKth(lo int, hi int, k int) int {
	nums := make([]int, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		nums = append(nums, i)
	}

	sort.Sort(Nums(nums))
	return nums[k-1]
}

type Nums []int

func (n Nums) Len() int      { return len(n) }
func (n Nums) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Nums) Less(i, j int) bool {
	p1, p2 := pow(n[i]), pow(n[j])
	if p1 < p2 {
		return true
	}

	if p1 > p2 {
		return false
	}

	return n[i] < n[j]
}

var mem = map[int]int{}

func pow(i int) int {
	if n, ok := mem[i]; ok {
		return n
	}

	if i == 1 {
		return 0
	}

	if i%2 == 0 {
		p := pow(i/2) + 1
		mem[i] = p
		return p
	}

	p := pow(i*3+1) + 1
	mem[i] = p
	return p
}

func TestIt(t *testing.T) {
	if getKth(12, 15, 2) != 13 {
		t.Error("error")
	}
	if getKth(1, 1000, 777) != 570 {
		t.Error("error")
	}
}
