package nearestexitfromentranceinmaze

import (
	"testing"
)

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	n := len(maze[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	queue := []*Cell{}

	appendCell := func(row, col int, step int) {
		if maze[row][col] == '.' && !visited[row][col] {
			queue = append(queue, &Cell{row, col, step})
		}
	}

	row, col := entrance[0], entrance[1]
	visited[row][col] = true
	if row > 0 {
		appendCell(row-1, col, 1)
	}
	if row < m-1 {
		appendCell(row+1, col, 1)
	}
	if col > 0 {
		appendCell(row, col-1, 1)
	}
	if col < n-1 {
		appendCell(row, col+1, 1)
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		if visited[cell.Row][cell.Col] {
			continue
		}

		if cell.Row == 0 || cell.Row == m-1 || cell.Col == 0 || cell.Col == n-1 {
			return cell.Step
		}

		visited[cell.Row][cell.Col] = true
		appendCell(cell.Row-1, cell.Col, 1+cell.Step)
		appendCell(cell.Row+1, cell.Col, 1+cell.Step)
		appendCell(cell.Row, cell.Col-1, 1+cell.Step)
		appendCell(cell.Row, cell.Col+1, 1+cell.Step)
	}

	return -1
}

type Cell struct {
	Row  int
	Col  int
	Step int
}

func TestNearestExit(t *testing.T) {
	maze := [][]byte{
		{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'},
	}
	entrance := []int{1, 2}
	ans := nearestExit(maze, entrance)

	if ans != 1 {
		t.Errorf("bad ans: %v", ans)
	}
}
