package answer

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/smallest-string-with-swaps/
// 这道题的基本思想是不要想着去模拟交换，而是根据间接相连的元素是可以任意交换的这一特点
// 对连通分量上的元素进行排序即可，求解连同分量可以使用图或者并查集
// 对连同分量组成的字符串排序：
// - 使用map记录每个连同分量的byte slice
// - 对map中的values进行排序
// - 遍历字符串，根据索引得到其连同分量，其值应当是连同分量的byte slice的相对索引上的byte
// - **技巧就是将byte slice作为消耗输入写入到ans中，通过消耗的方式写入队头，这样避免了记录索引、根据索引查找排序后的字符**
func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf := NewUF(len(s))
	for _, pair := range pairs {
		uf.union(pair[0], pair[1])
	}

	groups := map[int][]byte{}
	for i := range s {
		f := uf.find(i)
		groups[f] = append(groups[f], s[i])
	}

	for _, bytes := range groups {
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	}

	res := make([]byte, len(s))
	for i := range res {
		f := uf.find(i)
		res[i] = groups[f][0]
		groups[f] = groups[f][1:]
	}
	return string(res)
}

//标准版本的按秩求并的union find实现
type UF struct {
	id   []int
	rank []int
}

func NewUF(n int) *UF {
	id := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		rank[i] = 1
	}
	return &UF{
		id:   id,
		rank: rank,
	}
}

func (u *UF) connected(i, j int) bool {
	return u.find(i) == u.find(j)
}

func (u *UF) union(i, j int) {
	m, n := u.find(i), u.find(j)
	if m == n {
		return
	}
	if u.rank[m] < u.rank[n] {
		m, n = n, m
	}
	u.id[n] = m
	u.rank[m] += u.rank[n]
}

func (u *UF) find(i int) int {
	if i != u.id[i] {
		u.id[i] = u.find(u.id[i])
	}
	return u.id[i]
}

func TestIt(t *testing.T) {
	ss := []string{
		"dcab",
		"dcab",
	}
	pairss := [][][]int{
		{{0, 3}, {1, 2}},
		{{0, 3}, {1, 2}, {0, 2}},
	}

	expecteds := []string{
		"bacd",
		"abcd",
	}

	for i := range ss {
		s := ss[i]
		pairs := pairss[i]
		expected := expecteds[i]
		res := smallestStringWithSwaps(s, pairs)

		if res != expected {
			t.Errorf("expected %v, return %v", expected, res)
		}
	}
}
