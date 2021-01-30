package swim_in_rising_water

import "testing"

// https://leetcode-cn.com/problems/swim-in-rising-water/
// 使用union-find算法，只是复习了一下该算法，实际测试表明，使用该算法求解该问题效率不高
func swimInWater(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := NewUF(m * n)
	i := 0
	last := m*n - 1

	for {
		if uf.Connected(0, last) {
			break
		}
		i++
		for y := 0; y < m; y++ {
			for x := 0; x < n; x++ {
				if grid[y][x] <= i {
					curr := y*n + x
					if y > 0 && grid[y-1][x] <= i {
						top := curr - m
						uf.Union(curr, top)
					}
					if x > 0 && grid[y][x-1] <= i {
						left := curr - 1
						uf.Union(curr, left)
					}
				}
			}
		}
	}
	return i
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

type pos struct {
	x int
	y int
}

// 使用搜索算法，使用优先队列（最小堆）维护相邻但未标记的位置，每次取最小的值，标记并将相邻的未标记节点放入队列中，更新结果
// 直到遇到终点，如果可以立即到达，则保持结果不变，否则更新结果为终点的值
// 优化：为了避免重复插入，在插入时标记，而非取出时标记
func swimInWater1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	flood := make([][]bool, m)
	for i := 0; i < m; i++ {
		flood[i] = make([]bool, n)
	}

	heap := NewMinHeap(m * n)

	t := 0
	flood[0][0] = true
	heap.insert(grid[0][0], 0, 0)

	lastX, lastY := n-1, m-1
	for {
		if flood[lastY][lastX] {
			if grid[lastY][lastX] > t {
				t = grid[lastY][lastX]
			}
			break
		}
		val, x, y := heap.removeMin()
		if t < val {
			t = val
		}
		left, right, top, bottom := x-1, x+1, y-1, y+1
		if x > 0 && !flood[y][left] {
			flood[y][left] = true
			heap.insert(grid[y][left], left, y)
		}
		if right < n && !flood[y][right] {
			flood[y][right] = true
			heap.insert(grid[y][right], right, y)
		}
		if y > 0 && !flood[top][x] {
			flood[top][x] = true
			heap.insert(grid[top][x], x, top)
		}
		if bottom < m && !flood[bottom][x] {
			flood[bottom][x] = true
			heap.insert(grid[bottom][x], x, bottom)
		}
	}

	return t
}

type data struct {
	val, x, y int
}

type MinHeap struct {
	data      []*data
	cap, size int
}

func NewMinHeap(cap int) *MinHeap {
	return &MinHeap{
		data: make([]*data, cap+1, cap+1),
		cap:  cap,
		size: 0,
	}
}

func (m *MinHeap) insert(val, x, y int) {
	m.size++
	m.data[m.size] = &data{val, x, y}
	m.swim(m.size)
}

func (m *MinHeap) removeMin() (val, x, y int) {
	data := m.data[1]
	m.data[0] = m.data[m.size]
	m.size--
	m.sink(0)
	return data.val, data.x, data.y
}

func (m *MinHeap) swim(i int) {
	data := m.data[i]
	for i > 1 && m.data[i/2].val > data.val {
		m.data[i] = m.data[i/2]
		i = i / 2
	}
	m.data[i] = data
}

func (m *MinHeap) sink(i int) {
	data := m.data[i]
	for 2*i <= m.size {
		j := 2 * i
		if j < m.size && m.data[j].val > m.data[j+1].val {
			j++
		}
		if m.data[j].val >= data.val {
			break
		}
		m.data[i] = m.data[j]
		i = j
	}
	m.data[i] = data
}

func TestConn(t *testing.T) {
	var testCase1 [][]int
	testCase1 = append(testCase1, []int{0, 1, 2, 3, 4})
	testCase1 = append(testCase1, []int{24, 23, 22, 21, 5})
	testCase1 = append(testCase1, []int{12, 13, 14, 15, 16})
	testCase1 = append(testCase1, []int{11, 17, 18, 19, 20})
	testCase1 = append(testCase1, []int{10, 9, 8, 7, 6})

	result1 := swimInWater(testCase1)
	if result1 != 16 {
		t.Errorf("expected 16,but return %d", result1)
	}

	var testCase2 [][]int
	testCase2 = append(testCase2, []int{7, 8, 10, 13})
	testCase2 = append(testCase2, []int{2, 4, 15, 11})
	testCase2 = append(testCase2, []int{12, 1, 5, 6})
	testCase2 = append(testCase2, []int{9, 3, 0, 14})

	result2 := swimInWater(testCase2)
	if result2 != 14 {
		t.Errorf("expected 16,but return %d", result2)
	}
}
