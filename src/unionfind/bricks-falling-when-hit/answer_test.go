package answer

import "testing"

// https://leetcode-cn.com/problems/bricks-falling-when-hit/
// 建议参考题解https://leetcode-cn.com/problems/bricks-falling-when-hit/solution/803-da-zhuan-kuai-by-leetcode-r5kf/
//
// 这题很像游戏，但同时是一个特例，所有节点初始属于一个连同分量
// 最暴力的是，如果敲了一个有效砖块，每次求解连同分量，但效率最低
// 题解给出了一个有趣的解决办法，适合此类型的特例
func hitBricks(grid [][]int, hits [][]int) []int {
	rows, cols := len(grid), len(grid[0])
	rebuild := make([][]int, rows)

	// 第一步：复制grid但敲掉所有将敲掉的砖块
	for i := 0; i < rows; i++ {
		rebuild[i] = make([]int, cols)
		copy(rebuild[i], grid[i])
	}

	// 将敲击处置于0
	for _, hit := range hits {
		rebuild[hit[0]][hit[1]] = 0
	}

	size := rows * cols //top

	//第二步：构建并查集
	uf := NewUF(size + 1)
	index := func(row, col int) int { //这里使用x,y容易与传统笛卡尔坐标系混淆，因为x为row，y为col
		return row*cols + col
	}
	// 屋顶
	for i := 0; i < cols; i++ {
		if rebuild[0][i] == 1 {
			uf.Union(size, i)
		}
	}

	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//看其上、其左是否是砖块，是则相连
			if rebuild[i][j] == 1 {
				if rebuild[i-1][j] == 1 {
					uf.Union(index(i-1, j), index(i, j))
				}
				if j > 0 && rebuild[i][j-1] == 1 {
					uf.Union(index(i, j-1), index(i, j))
				}
			}
		}
	}

	//第三步：逆序补回砖块
	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		row, col := hits[i][0], hits[i][1]
		if grid[row][col] == 0 {
			continue
		}

		rebuild[row][col] = 1

		prevSize := uf.Size(size)
		if row == 0 {
			uf.Union(size, index(row, col))
		} else if rebuild[row-1][col] == 1 {
			uf.Union(index(row-1, col), index(row, col))
		}

		if row < rows-1 && rebuild[row+1][col] == 1 {
			uf.Union(index(row+1, col), index(row, col))
		}

		if col > 0 && rebuild[row][col-1] == 1 {
			uf.Union(index(row, col-1), index(row, col))
		}

		if col < cols-1 && rebuild[row][col+1] == 1 {
			uf.Union(index(row, col+1), index(row, col))
		}

		currSize := uf.Size(size)

		remove := currSize - prevSize - 1
		if remove < 0 {
			remove = 0
		}

		ans[i] = remove

	}

	return ans
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
	grids := [][][]int{
		{{1}, {1}, {1}, {1}, {1}},
		{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}},
	}
	hitss := [][][]int{
		{{3, 0}, {4, 0}, {1, 0}, {2, 0}, {0, 0}},
		{{0, 2}, {2, 0}, {0, 1}, {1, 2}},
	}
	expecteds := [][]int{
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0},
	}

	for i, grid := range grids {
		hits := hitss[i]
		expected := expecteds[i]
		res := hitBricks(grid, hits)
		for j := 0; j < len(res); j++ {
			if res[j] != expected[j] {
				t.Errorf("expected %v, return %v", expected, res)
				break
			}
		}
	}
}
