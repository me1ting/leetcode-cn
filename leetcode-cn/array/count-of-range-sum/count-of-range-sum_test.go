package count_of_range_sum

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/count-of-range-sum
// 自己是想不出来的,理解题解即可
// https://leetcode-cn.com/problems/count-of-range-sum/solution/qu-jian-he-de-ge-shu-by-leetcode-solution/
// 核心在于前缀和和归并排序
// 使用归并排序来计算区间和位于目标范围的个数
// 注意的是，不能先排序再计算差值，因为顺序很重要，而归并排序的特色在于：
//   两个值之间的差值只需要计算一次，保证计算时的前后顺序（谁减谁）
func countRangeSum(nums []int, lower int, upper int) int {
	var s int64 = 0
	sum := make([]int64, len(nums)+1)
	for i, n := range nums {
		s += int64(n)
		sum[i+1] = s
	}

	return mergeCount(sum, int64(lower), int64(upper))
}

func mergeCount(part []int64, lower int64, upper int64) int {
	if len(part) == 1 {
		return 0
	}
	mid := (len(part) - 1) / 2
	ret := mergeCount(part[:mid+1], lower, upper) + mergeCount(part[mid+1:], lower, upper)

	for i, l, r := 0, mid+1, mid+1; i <= mid; i++ {
		for l < len(part) && part[l]-part[i] < lower {
			l++
		}
		for r < len(part) && part[r]-part[i] <= upper {
			r++
		}
		ret += r - l
	}

	//mergeSort
	sorted:=make([]int64,len(part))
	p1,p2,p:=0,mid+1,0
	for p1<=mid&&p2<len(part)  {
		if part[p1]<part[p2] {
			sorted[p] = part[p1]
			p1++
		}else{
			sorted[p]=part[p2]
			p2++
		}
		p++
	}

	for ;p1<=mid;p1++ {
		sorted[p] = part[p1]
		p++
	}
	for ;p2<len(part);p2++ {
		sorted[p] = part[p2]
		p++
	}

	copy(part,sorted)

	return ret
}

func TestCount(t *testing.T) {
	nums:=[]int{0}
	lower:=-2
	upper:=2

	fmt.Println(countRangeSum(nums,lower,upper))
}
