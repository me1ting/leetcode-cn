package hj43migong_test

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func main() {
	inputs := readInputs()
	path := Answer(inputs)

	for _, p := range path {
		fmt.Printf("(%v,%v)", p.Y, p.X)
	}
}

func readInputs() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	metaLine := strings.TrimSpace(scanner.Text())
	pieces := strings.Split(metaLine, " ")
	R, _ := strconv.Atoi(pieces[0])
	//C, _ := strconv.Atoi(pieces[1])

	rows := [][]int{}
	for i := 0; i < R; i++ {
		scanner.Scan()
		row := []int{}
		line := strings.TrimSpace(scanner.Text())
		pieces := strings.Split(line, " ")
		for _, p := range pieces {
			n, _ := strconv.Atoi(p)
			row = append(row, n)
		}
		rows = append(rows, row)
	}
	return rows
}

func Answer(rows [][]int) []*point {
	R, C := len(rows), len(rows[0])

	graphs := make([][]*point, R)
	for i := 0; i < R; i++ {
		graphs[i] = make([]*point, C)
	}

	h := &PointSlice{}
	graphs[0][0] = &point{0, 0, 0, nil, false}
	heap.Push(h, graphs[0][0])

	var update = func(x, y int, from *point) {
		if x < 0 || x >= C || y < 0 || y >= R {
			return
		}
		if rows[y][x] == 0 {
			p := graphs[y][x]
			i := from.Indegree + 1
			if p == nil {
				p = &point{x, y, i, from, false}
				graphs[y][x] = p
				heap.Push(h, p)
			} else {
				if !p.visited && p.Indegree > i {
					p.Indegree = i
					heap.Push(h, p)
				}
			}
		}
	}

	outX, outY := C-1, R-1

	for h.Len() > 0 {
		p := heap.Pop(h).(*point)
		if !p.visited {
			p.visited = true
			if p.X == outX && p.Y == outY {
				break
			}
			update(p.X-1, p.Y, p)
			update(p.X+1, p.Y, p)
			update(p.X, p.Y-1, p)
			update(p.X, p.Y+1, p)
		}
	}

	path := []*point{}
	for p := graphs[outY][outX]; p != nil; p = p.from {
		path = append(path, p)
	}

	for i, j := 0, len(path)-1; i < j; {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}

	return path
}

type point struct {
	X        int
	Y        int
	Indegree int
	from     *point
	visited  bool
}

type PointSlice []*point

// Len implements heap.Interface.
func (p *PointSlice) Len() int {
	return len(*p)
}

// Less implements heap.Interface.
func (p *PointSlice) Less(i int, j int) bool {
	return (*p)[i].Indegree < (*p)[j].Indegree
}

// Pop implements heap.Interface.
func (p *PointSlice) Pop() any {
	item := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return item
}

// Push implements heap.Interface.
func (p *PointSlice) Push(x any) {
	*p = append(*p, x.(*point))
}

// Swap implements heap.Interface.
func (p *PointSlice) Swap(i int, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

var _ heap.Interface = &PointSlice{}

func TestAnswer(t *testing.T) {
	inputs := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}

	Answer(inputs)
}
