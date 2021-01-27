package answer

// https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/
// 官方题解讲得很清晰：
// https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/solution/bao-zheng-tu-ke-wan-quan-bian-li-by-leet-mtrw/
//
// 公共边优先于私有边，因为使用私有边需要两条，而使用公共边只需要一条

func maxNumEdgesToRemove(n int, edges [][]int) int {
	uf := NewUF(n + 1)
	removed := 0
	for _, edge := range edges {
		if edge[0] == 3 {
			if !uf.Union(edge[1], edge[2]) {
				removed++
			}
		}
	}

	cp := uf.Clone()
	for _, edge := range edges {
		if edge[0] == 1 {
			if !cp.Union(edge[1], edge[2]) {
				removed++
			}
		}
	}

	if cp.setCount != 2 { //使用2使用 位置0并没有被用到 因此多一个连通分量
		return -1
	}
	for _, edge := range edges {
		if edge[0] == 2 {
			if !uf.Union(edge[1], edge[2]) {
				removed++
			}
		}
	}
	if uf.setCount != 2 {
		return -1
	}

	return removed
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

func (u *UF) Clone() *UF {
	n := len(u.size)
	id := make([]int, n)
	size := make([]int, n)
	copy(id, u.id)
	copy(size, u.size)

	return &UF{
		id:       id,
		size:     size,
		setCount: u.setCount,
	}
}
