package answer

import "testing"

// https://leetcode-cn.com/problems/regions-cut-by-slashes/
// 此题明显是连通分量的问题，但是如何将题目映射到连通性上面是很难直觉得出的
// 参考题解，将1X1区块划分为4部分（因为只有\,/, 三种情况）：
//
// 0\1/2
// 0/3\2
//
// 根据不同的情况进行合并操作。注意的是，对于相连的区块的相连部分也需要进行合并

func regionsBySlashes(grid []string) int {
	n := len(grid)
	uf := NewUF(n * n * 4)

	for i, row := range grid {
		for j, cell := range row {
			index := (i*n + j) * 4
			switch cell {
			case '\\':
				uf.Union(index, index+3)
				uf.Union(index+1, index+2)
			case '/':
				uf.Union(index, index+1)
				uf.Union(index+2, index+3)
			case ' ':
				uf.Union(index, index+1)
				uf.Union(index, index+2)
				uf.Union(index, index+3)
			}
			if i > 0 {
				uf.Union(index+1, index-n*4+3)
			}
			if j > 0 {
				uf.Union(index, index-4+2)
			}
		}
	}

	return uf.setCount
}

type UF struct {
	id       []int
	size     []int
	setCount int
}

func NewUF(n int) *UF {
	id := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		size[i] = 1
	}
	return &UF{
		id:       id,
		size:     size,
		setCount: n,
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
	if u.size[m] < u.size[n] {
		m, n = n, m
	}
	u.id[n] = m
	u.size[m] += u.size[n]
	u.setCount--

	return true
}

func (u *UF) Find(i int) int {
	if i != u.id[i] {
		u.id[i] = u.Find(u.id[i])
	}
	return u.id[i]
}

func (u *UF) SetCount() int {
	return u.setCount
}

func TestRegionsBySlashes(t *testing.T) {
	testcase := []string{" /", "/ "}
	excepted := 2
	val := regionsBySlashes(testcase)

	if val != excepted {
		t.Errorf("expected %v,reutrn %v", excepted, val)
	}
}
