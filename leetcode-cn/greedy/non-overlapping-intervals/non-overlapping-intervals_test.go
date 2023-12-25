package non_overlapping_intervals

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/non-overlapping-intervals/
// 这道题的思想其实和 https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons 很类似
// 解法：
// 将所有区间按照左位置排序：
// 遍历排序后的区间，可能出现以下情况：
// 1. 覆盖，保留后一个区间
// 2. 交叉，保留前一个区间
// 3. 不相接，保留两者
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Sort(Intervals(intervals))
	removed := 0
	curr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if next[0] < curr[1] {
			if next[1] <= curr[1] {
				curr = next
			}
			removed++
		} else {
			curr = next
		}
	}
	return removed
}

type Intervals [][]int

func (a Intervals) Len() int           { return len(a) }
func (a Intervals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Intervals) Less(i, j int) bool { return a[i][0] < a[j][0] }

func TestIt(t *testing.T) {
	testcase := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	if eraseOverlapIntervals(testcase) != 1 {
		t.Errorf("error")
	}
}
