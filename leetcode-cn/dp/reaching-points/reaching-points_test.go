package reaching_points

import "testing"

// https://leetcode-cn.com/problems/reaching-points/
// 使用深度优先搜索（用到了动态规划）
// 实际测试，这是错误的做法，对于很大的tx,ty，最终会导致OOM
func reachingPoints1(sx int, sy int, tx int, ty int) bool {
	mark := make([][]bool, ty+1)
	for i := 0; i <= ty; i++ {
		mark[i] = make([]bool, tx+1)
	}
	if sx > tx || sy > ty {
		return false
	}

	if sx == tx && sy == ty {
		return true
	}

	stack := NewStack(max(tx, ty))
	stack.push(sx, sy)
	for stack.size > 0 {
		x, y := stack.pop()
		if x == tx && y == ty {
			return true
		}
		xy := x + y

		if (x == tx && xy == ty) || (xy == tx && y == ty) {
			return true
		}
		if xy <= tx && !mark[y][xy] {
			mark[y][xy] = true
			stack.push(xy, y)
		}
		if xy <= ty && !mark[xy][x] {
			mark[xy][x] = true
			stack.push(x, xy)
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type data struct {
	x, y int
}

type stack struct {
	data []*data
	size int
	cap  int
}

func NewStack(cap int) *stack {
	return &stack{
		data: make([]*data, cap),
		size: 0,
		cap:  cap,
	}
}

func (s *stack) push(x, y int) {
	if s.size == s.cap {
		newData := make([]*data, s.cap*2)
		copy(newData, s.data)
		s.cap *= 2
	}

	s.data[s.size] = &data{x, y}
	s.size++
}

func (s *stack) pop() (x, y int) {
	s.size--
	data := s.data[s.size]
	return data.x, data.y
}

// 其实这题是考察对动态规划、回溯法的理解
// 回溯法：由终点倒推起点
// 对于此题，每个点它的前一个节点是固定的
// 对于此题，当sx==1||sy==1时，还是会产生超时的问题

func reachingPoints2(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx == sx && ty == sy {
			return true
		}
		if tx > ty {
			tx -= ty
		} else {
			ty -= tx
		}
	}
	return false
}

// 如果一个对是 (a,b)，那么它之前必然是 (a,b−a)或者 (a−b,b)
// 只有当a′ 变成 a mod b时，才会发生交换
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if tx==ty{
			break
		}
		if tx>ty {
			if ty>sy {
				tx%=ty
			}else{
				return (tx-sx)%ty == 0
			}
		}else{
			if tx>sx {
				ty%=tx
			}else{
				return (ty-sy)%tx == 0
			}
		}
	}

	return tx==sx&&ty==sy
}

func TestReachingPoints(t *testing.T) {
	if reachingPoints(1, 1, 3, 5) != true {
		t.Errorf("expected true,but return false")
	}
	if reachingPoints(9, 10, 9, 19) != true {
		t.Errorf("expected true,but return false")
	}
}
