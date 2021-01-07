package answer

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/largest-number/submissions/
// 对数字转换为字符串后进行排序
// 开头数字大的在前面，一个消耗思维的点是：
// 如何比较55xxx，55的大小？
//
// 不难，但实现起来繁琐，要考虑边界情况
func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	sort.Sort(Nums(strs))

	// 边界情况，只包含0
	if strs[0][0] == '0' {
		return "0" //因为返回"00"是错误的
	}

	return strings.Join(strs, "")
}

type Nums []string

func (n Nums) Len() int      { return len(n) }
func (n Nums) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Nums) Less(i, j int) bool {
	//逆序
	return !less(n[i], n[j])
}

func less(s1, s2 string) bool {
	cmpLen := min(len(s1), len(s2))
	for i := 0; i < cmpLen; i++ {
		if s1[i] < s2[i] {
			return true
		}
		if s1[i] > s2[i] {
			return false
		}
	}

	//55xxx,55比较大小，等价于：
	//55xxx55
	//5555xxx
	//这两个数比较大小
	s3 := s1[cmpLen:] + s2
	s4 := s2[cmpLen:] + s1
	for i := 0; i < len(s3); i++ {
		if s3[i] < s4[i] {
			return true
		}
		if s3[i] > s4[i] {
			return false
		}
	}

	return true
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
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
