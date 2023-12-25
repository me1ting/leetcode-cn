package main

import (
	"fmt"
	"testing"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	n := len(grid)
	shortestPaths := make([][]int, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		shortestPaths[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}
	shortestPaths[0][0] = 1

	stack := []*Vertex{}

	update := func(r, c int, path int) {
		if r < 0 || r >= n || c < 0 || c >= n {
			return
		}
		if grid[r][c] == 1 {
			return
		}
		if visited[r][c] {
			return
		}
		if shortestPaths[r][c] == 0 {
			shortestPaths[r][c] = path
			stack = append(stack, &Vertex{r, c})
		}
	}

	stack = append(stack, &Vertex{0, 0})
	for len(stack) > 0 {
		v := stack[0]
		row := v.row
		col := v.col
		visited[row][col] = true
		path := shortestPaths[row][col] + 1
		for c := col - 1; c <= col+1; c++ {
			for r := row - 1; r <= row+1; r++ {
				update(r, c, path)
			}
		}
		stack = stack[1:]
	}

	if !visited[n-1][n-1] {
		return -1
	}
	return shortestPaths[n-1][n-1]
}

type Vertex struct {
	row int
	col int
}

func (v *Vertex) String() string {
	return fmt.Sprintf(`{%d, %d}`, v.row, v.col)
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	gridList := [][][]int{
		{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}},
		{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}},
	}

	ansList := []int{
		4,
		-1,
	}

	for i, grid := range gridList {
		ans := ansList[i]
		res := shortestPathBinaryMatrix(grid)

		if res != ans {
			t.Errorf("expected %v, got %v", ans, res)
		}
	}
}
