package answer

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/largest-number/
// 要获得最大的整数，一个很简单的想法是“高个子顶前面”：
// 9... > 8... > 7... > ... > 0
// 98... > 97... > 96... >... > 90...
// 对于两个数字，我们从高位依次比较每一位，如果其中某一位不相等，那么两个数字的大小比较结果等于该位的比较结果。
// 但是存在特殊情况，当一个数字是另一个数字前缀时，如55和55...，该怎么比较呢？
// 这可以通过一种简单的办法来比较：cmp(55,55...) == cmp(55 55...,55... 55)。
//
// 为了按十进制位从高比较两个数字，并在排序过程中进行多次比较，最好的办法是缓存每个数字的十进制位数组，
// 而Go的string正好满足我们的需求。
//
// 最后的返回结果存在一种特殊情况需要处理："0...0"，返回"0"。
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	sort.Sort(sort.Reverse(Nums(strs))) // Go的排序为从小到大，使用逆序转换为从大到小

	// 如果结果是"0...0"，返回"0"
	if strs[0][0] == '0' {
		return "0"
	}

	return strings.Join(strs, "")
}

type Nums []string

func (n Nums) Len() int      { return len(n) }
func (n Nums) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Nums) Less(i, j int) bool {
	// 子过程，实现了cmp("55"+"55..."","55..."+"55")
	// 子过程仅假设len(s1) <= len(s2)
	less := func(s1, s2 string) bool {
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				return s1[i] < s2[i]
			}
		}
		for i, j := 0, len(s1); j < len(s2); i, j = i+1, j+1 {
			if s2[i] != s2[j] {
				return s2[i] < s2[j]
			}
		}
		for i, j := len(s2)-len(s1), 0; i < len(s2); i, j = i+1, j+1 {
			if s2[i] != s1[j] {
				return s2[i] < s1[j]
			}
		}
		return false
	}

	s1, s2 := n[i], n[j]
	if len(s1) <= len(s2) {
		return less(s1, s2)
	}
	return !less(s2, s1)
}

func TestIt(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	expected := "9534330"
	res := largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}

	nums = []int{10, 2}
	expected = "210"
	res = largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}

	nums = []int{1}
	expected = "1"
	res = largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}

	nums = []int{111311, 1113}
	expected = "1113111311"
	res = largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}

	nums = []int{1113, 111311}
	//1113 111311
	//111311 1113
	expected = "1113111311"
	res = largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}

	nums = []int{8308, 8308, 830}
	expected = "83088308830"
	res = largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}

	nums = []int{0, 0}
	expected = "0"
	res = largestNumber(nums)
	if res != expected {
		t.Errorf("expected %v,return %v", expected, res)
	}
}
