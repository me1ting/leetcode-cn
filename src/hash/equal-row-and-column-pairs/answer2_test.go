package answer

import "testing"

func equalPairs2(grid [][]int) int {
	n := len(grid[0])
	count := map[int]entryList{}
	for _, row := range grid {
		h := hash(row)
		list := count[h]
		if list == nil {
			count[h] = entryList{&entry{row, 1}}
		} else {
			list.setOrAdd(row, 1)
		}
	}

	ans := 0

	for i := 0; i < n; i++ {
		col := make([]int, n)
		for j := 0; j < n; j++ {
			col[j] = grid[j][i]
		}
		h := hash(col)
		list := count[h]
		if list != nil {
			ans += list.get(col)
		}
	}
	return ans
}

func hash(line []int) int {
	h := 0
	for _, num := range line {
		h = 31*h + num
	}
	return h
}

type entry struct {
	line []int
	val  int
}

type entryList []*entry

func (el *entryList) get(line []int) int {
	for i := 0; i < len(*el); i++ {
		if equals((*el)[i].line, line) {
			return (*el)[i].val
		}
	}
	return 0
}

func (el *entryList) setOrAdd(line []int, val int) {
	for i := 0; i < len(*el); i++ {
		if equals((*el)[i].line, line) {
			(*el)[i].val += val
			return
		}
	}

	*el = append(*el, &entry{line, val})
}

func equals(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestEqualPairs2(t *testing.T) {
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	rightAns := 3

	ans := equalPairs2(grid)
	if ans != rightAns {
		t.Errorf("expected %v, but got %v", rightAns, ans)
	}
}
