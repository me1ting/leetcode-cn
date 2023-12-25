package answer

import "testing"

// https://leetcode-cn.com/problems/couples-holding-hands/solution/
// 此题的核心算法是：
// 将一队情侣视为点，将位置坐错视为点之间的连接，就可以将当前的状态视为一幅图
// 对于图中的连通分量，需要交换其大小m减去1次，n个连通分量的交换次数求和：
// sum(m_n-1) = sum(m_n)-sum(1) = N - n
func minSwapsCouples(row []int) int {
	n := len(row)
	coupleN := n / 2

	uf := NewUF(coupleN)
	for i := 0; i < n; i += 2 {
		uf.Union(row[i]/2, row[i+1]/2)
	}

	return coupleN - uf.Count()
}

//标准版UF
type UF struct {
	id    []int
	rank  []int
	count int
}

func NewUF(n int) *UF {
	id := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		rank[i] = 1
	}
	return &UF{
		id:    id,
		rank:  rank,
		count: n,
	}
}

func (u *UF) Connected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

func (u *UF) Union(i, j int) bool {
	m, n := u.Find(i), u.Find(j)
	if m == n {
		return false
	}
	if u.rank[m] < u.rank[n] {
		m, n = n, m
	}
	u.id[n] = m
	u.rank[m] += u.rank[n]
	u.count--

	return true
}

func (u *UF) Find(i int) int {
	if i != u.id[i] {
		u.id[i] = u.Find(u.id[i])
	}
	return u.id[i]
}

func (u *UF) Count() int {
	return u.count
}

func TestMinSwapsCouples(t *testing.T) {
	rows := [][]int{
		{0, 2, 1, 3},
		{3, 2, 0, 1},
	}
	expecteds := []int{
		1, 0,
	}

	for i, row := range rows {
		expected := expecteds[i]
		val := minSwapsCouples(row)
		if val != expected {
			t.Errorf("return %v, expected %v", val, expected)
		}
	}
}
