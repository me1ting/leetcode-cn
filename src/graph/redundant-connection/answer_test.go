package answer

// https://leetcode-cn.com/problems/redundant-connection
// 此题是求解图中的环，答案要求返回构成环的最后一条边
// 根据要求的答案，使用并查集很适合
func findRedundantConnection(edges [][]int) []int {
	uf := NewUF(len(edges) + 1) // 顶点个数为len，但又因为顶点从1开始
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		fu, fv := uf.Find(u), uf.Find(v)
		if fu == fv {
			return []int{u, v}
		}
		uf.Union(u, v)
	}
	return nil
}

type UF struct {
	id   []int
	size []int
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
		size: rank,
	}
}

func (u *UF) Connected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}

func (u *UF) Union(i, j int) {
	m, n := u.Find(i), u.Find(j)
	if m == n {
		return
	}
	if u.size[m] < u.size[n] {
		m, n = n, m
	}
	u.id[n] = m
	u.size[m] += u.size[n]
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
