package numberofwaystoarriveatdestination_test

import (
	"container/heap"
	"math"
	"testing"
)

// https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination
// 思路：单源最短优先路径，使用优先队列的Dijkstra算法
func countPaths(n int, roads [][]int) int {
	const mod = 1e9 + 7
	g := make([][]Edge, n)
	for _, r := range roads {
		u, v, cost := r[0], r[1], r[2]
		g[u] = append(g[u], Edge{v, cost})
		g[v] = append(g[v], Edge{u, cost})
	}

	dis := make([]int64, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt64
	}

	ways := make([]int64, n)
	ways[0] = 1

	q := PriorityQueue{{0, 0}}
	heap.Init(&q)

	for len(q) > 0 {
		entry := heap.Pop(&q).(*Entry)
		u, d := entry.u, entry.dis
		if d > dis[u] {
			continue
		}
		for _, edge := range g[u] {
			v, c := edge.v, int64(edge.cost)
			disV := d + c
			if disV < dis[v] {
				dis[v] = disV
				ways[v] = ways[u]
				heap.Push(&q, &Entry{v, disV})
			} else if disV == dis[v] {
				ways[v] = (ways[u] + ways[v]) % mod
			}
		}
	}

	return int(ways[n-1])
}

type Edge struct {
	v    int
	cost int
}

type Entry struct {
	u   int
	dis int64
}

type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dis < pq[j].dis
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Entry))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return x
}

var _ heap.Interface = &PriorityQueue{}

func TestIt(t *testing.T) {
	n := 7
	roads := [][]int{{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3},
		{3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2}}
	if countPaths(n, roads) != 4 {
		t.Fatal("bad answer")
	}
}
