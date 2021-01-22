package answer

import (
	"math"
	"sort"
	"testing"
)

type unionFind struct {
	parent, size []int
	setCount     int // 当前连通分量数目，用于后续的特殊目的
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	// 将边在edges中的索引值记录在边的信息中
	for i, e := range edges {
		edges[i] = append(e, i)
	}
	// 将所有边按照权进行排序
	sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

	// 计算最小生成树，允许跳过某一边（根据索引），允许自定义unionfind（可以预先确定一些边）
	// 返回MST的大小
	calcMST := func(uf *unionFind, ignoreID int) (mstValue int) {
		for i, e := range edges {
			if i != ignoreID && uf.union(e[0], e[1]) {
				mstValue += e[2]
			}
		}
		if uf.setCount > 1 {
			return math.MaxInt64
		}
		return
	}

	//计算最小生成树
	mstValue := calcMST(newUnionFind(n), -1)

	var keyEdges, pseudokeyEdges []int
	for i, e := range edges {
		// 是否为关键边
		if calcMST(newUnionFind(n), i) > mstValue {
			keyEdges = append(keyEdges, e[3])
			continue
		}

		// 是否为伪关键边
		uf := newUnionFind(n)
		uf.union(e[0], e[1])
		if e[2]+calcMST(uf, i) == mstValue {
			pseudokeyEdges = append(pseudokeyEdges, e[3])
		}
	}

	return [][]int{keyEdges, pseudokeyEdges}
}

func TestIt(t *testing.T) {
	n := 4
	edges := [][]int{
		{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1},
	}

	expected := [][]int{{}, {0, 1, 2, 3}}
	val := findCriticalAndPseudoCriticalEdges(n, edges)

	if len(val) != len(expected) {
		t.Errorf("expected %v,return %v", expected, val)
		return
	}

	for i := 0; i < len(val); i++ {
		if len(val[i]) != len(expected[i]) {
			t.Errorf("expected %v,return %v", expected, val)
			return
		}
		for j := 0; j < len(val[i]); j++ {
			if val[i][j] != expected[i][j] {
				t.Errorf("expected %v,return %v", expected, val)
				return
			}
		}
	}

}
