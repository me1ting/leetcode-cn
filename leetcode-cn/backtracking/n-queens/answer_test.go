package nqueens

import (
	"log"
	"testing"
)

// https://leetcode.cn/problems/n-queens
// 还是以递归的思想来解决这个问题，使用如下方式表示一个4*4的棋盘：
// y
// 3
// 2
// 1
// 0 1 2 3 x
//
// 使用`states`变量记录了横向、纵向、斜度1，斜度-1方向上是否已经有棋子。
// 我们从第0行开始扫描行中的每格，如果能够放下棋子，那就放下棋子然后扫描下一行，直到扫描完所有行，记录结果。
func solveNQueens(n int) [][]string {
	res := [][]string{}

	queens := make([]int, n)                                                   //记录每行的皇后的位置
	rowStates, colStates := make([]bool, n), make([]bool, n)                   //横向、纵向
	incline1States, incline2States := make([]bool, n*2-1), make([]bool, n*2-1) //斜度1方向，斜度-1方向

	check := func(x, y int) bool {
		return !(rowStates[y] || colStates[x] || incline1States[y-x+n-1] || incline2States[y+x])
	}

	place := func(x, y int) {
		queens[y] = x
		rowStates[y] = true
		colStates[x] = true
		incline1States[y-x+n-1] = true
		incline2States[y+x] = true
	}

	remove := func(x, y int) {
		rowStates[y] = false
		colStates[x] = false
		incline1States[y-x+n-1] = false
		incline2States[y+x] = false
	}

	var scan func(y int) // 扫描第 y 行
	scan = func(y int) {
		if y == n {
			res = append(res, queensToStrings(queens))
			return
		}
		for x := range n {
			if check(x, y) {
				place(x, y)
				scan(y + 1)
				remove(x, y)
			}
		}
	}

	scan(0)

	return res
}

func queensToStrings(queens []int) []string {
	res := []string{}
	bytes := make([]byte, len(queens))
	for i := 0; i < len(bytes); i++ {
		bytes[i] = '.'
	}
	for _, q := range queens {
		bytes[q] = 'Q'
		res = append(res, string(bytes))
		bytes[q] = '.'
	}

	return res
}

func TestSolveNQueens(t *testing.T) {
	log.Println(solveNQueens(4))
}
