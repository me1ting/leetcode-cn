package main

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/min-cost-to-connect-all-points
// 最小生成树
// 使用暴力法计算所有边的权（曼哈顿）需要O(n^2)的时间复杂度
// TODO 理解优化算法并进行优化
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	edges := make([]*edge, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]

			edges = append(edges, &edge{
				w:   i,
				v:   j,
				dis: distance(x1, y1, x2, y2),
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis < edges[j].dis
	})

	uf := NewUF(n)

	i := 0
	cost := 0
	for _, edge := range edges {
		if i == n-1 {
			break
		}
		w, v := edge.w, edge.v
		if uf.Connected(w, v) {
			continue
		}
		uf.Union(w, v)
		cost += edge.dis
		i++
	}
	return cost
}

type edge struct {
	w, v int
	dis  int
}

func distance(x1, y1, x2, y2 int) int {
	dis := 0
	if x1 < x2 {
		dis += x2 - x1
	} else {
		dis += x1 - x2
	}

	if y1 < y2 {
		dis += y2 - y1
	} else {
		dis += y1 - y2
	}
	return dis
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

func TestIt(t *testing.T) {
	points := [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}
	expected := 20
	result := minCostConnectPoints(points)
	if result != expected {
		t.Errorf("expected %v, return %v", expected, result)
	}
}
