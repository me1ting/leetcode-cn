package main

import (
	"container/heap"
	"math"
	"testing"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
		for j := 0; j < n; j++ {
			paths[i][j] = math.MaxInt
		}
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	pq := PQ([]*Record{})

	check := func(row, col int, way int) {
		if row == m || col == n || visited[row][col] {
			return
		}
		newPath := way + grid[row][col]
		if newPath < paths[row][col] {
			paths[row][col] = newPath
			heap.Push(&pq, &Record{row, col, newPath})
		}
	}

	update := func(row, col int) {
		visited[row][col] = true
		p := paths[row][col]
		check(row+1, col, p)
		check(row, col+1, p)
	}

	paths[0][0] = grid[0][0]
	heap.Push(&pq, &Record{0, 0, paths[0][0]})

	for pq.Len() > 0 {
		v := heap.Pop(&pq).(*Record)
		if visited[v.row][v.col] {
			continue
		}
		update(v.row, v.col)
	}
	return paths[m-1][n-1]
}

type PQ []*Record

type Record struct {
	row  int
	col  int
	path int
}

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	ri, rj := pq[i], pq[j]
	return ri.path < rj.path
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x any) {
	v := x.(*Record)
	*pq = append(*pq, v)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func TestMinPathSum(t *testing.T) {
	gridList := [][][]int{
		{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		{{1, 2, 3}, {4, 5, 6}},
	}
	rightAnsList := []int{7, 12}

	for i, grid := range gridList {
		rightAns := rightAnsList[i]

		res := minPathSum(grid)

		if res != rightAns {
			t.Errorf("expected %v, returns %v", rightAns, res)
		}
	}
}
