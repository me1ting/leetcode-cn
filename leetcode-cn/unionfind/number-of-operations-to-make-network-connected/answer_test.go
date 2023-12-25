package answer

// https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected/
// 形成环的边是多余的，每一个独立的连通分量需要额外一条边
// 使用并查集，统计多余的边，统计连通分量大小
func makeConnected(n int, connections [][]int) int {
	uf := NewUF(n)
	edgeN := 0
	for _, edge := range connections {
		if !uf.Union(edge[0], edge[1]) {
			edgeN++
		}
	}

	if edgeN >= uf.setCount-1 {
		return uf.setCount - 1
	}
	return -1
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

// Union方法返回值表示是否进行了合并操作，这个技巧在某些场景很有用
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

func (u *UF) Size(i int) int {
	return u.size[u.Find(i)]
}
