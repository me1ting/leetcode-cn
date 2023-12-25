package answer

import "sort"

// https://leetcode-cn.com/problems/path-with-minimum-effort/
// 参考官方题解
// 将地图转换为图，差的绝对值即为两点之间的边，问题通过图问题得到解决
// 可以使用并查集，将所有边排序，使用边按排序顺序进行合并，直到两点为同一连同分量
func minimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])

	type edge struct {
		w int
		v int
		e int
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	edges := make([]*edge, 0, m*(n-1)+n*(m-1))
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				edges = append(edges, &edge{
					w: i*n + j - 1,
					v: i*n + j,
					e: abs(heights[i][j] - heights[i][j-1]),
				})
			}
			if i > 0 {
				edges = append(edges, &edge{
					w: (i-1)*n + j,
					v: i*n + j,
					e: abs(heights[i][j] - heights[i-1][j]),
				})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].e < edges[j].e
	})

	uf := NewUF(m * n)
	end := m*n - 1

	for _, edge := range edges {
		uf.Union(edge.w, edge.v)
		if uf.Connected(0, end) {
			return edge.e
		}
	}

	return 0

}

type UF struct {
	id   []int
	size []int
}

func NewUF(n int) *UF {
	id := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		size[i] = 1
	}
	return &UF{
		id:   id,
		size: size,
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

	return true
}

func (u *UF) Find(i int) int {
	if i != u.id[i] {
		u.id[i] = u.Find(u.id[i])
	}
	return u.id[i]
}
