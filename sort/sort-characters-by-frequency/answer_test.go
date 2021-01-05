package answer

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/sort-characters-by-frequency/
// 桶、hashmap、堆...只要实现计数、排序功能的都可以
func frequencySort(s string) string {
	counts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	countArray := make([]*Count, 0, len(counts))
	for c, n := range counts {
		countArray = append(countArray, &Count{c, n})
	}

	sort.Sort(CountArray(countArray))

	res := make([]byte, 0, len(s))

	for _, count := range countArray {
		for i := 0; i < count.n; i++ {
			res = append(res, count.c)
		}
	}

	return string(res)
}

type Count struct {
	c byte
	n int
}

type CountArray []*Count

func (ca CountArray) Len() int           { return len(ca) }
func (ca CountArray) Less(i, j int) bool { return ca[i].n > ca[j].n }
func (ca CountArray) Swap(i, j int)      { ca[i], ca[j] = ca[j], ca[i] }

func TestIt(t *testing.T) {
	result := frequencySort("tree")
	if result != "eert" && result != "eetr" {
		t.Errorf("error: expected %v, but %v.", "eert", result)
	}

	result = frequencySort("cccaab")
	if result != "cccaab" {
		t.Errorf("error: expected %v, but %v.", "cccaab", result)
	}
}
