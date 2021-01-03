package answer

import "testing"

// https://leetcode-cn.com/problems/daily-temperatures/
// 使用栈记录单调递减的值的索引(要点)
// 对数组进行逆向遍历，每次出栈值小于等于当前值的索引，记录下ans，入栈当前值的索引
func dailyTemperatures(T []int) []int {
	n := len(T)
	stack := NewStack(n)
	ans := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for stack.size > 0 && T[stack.top()] <= T[i] {
			stack.pop()
		}

		if stack.size > 0 {
			ans[i] = stack.top() - i
		} else {
			ans[i] = 0
		}
		stack.push(i)
	}

	return ans
}

type IntStack struct {
	data []int
	size int
	cap  int
}

func NewStack(cap int) *IntStack {
	return &IntStack{
		data: make([]int, cap),
		size: 0,
		cap:  cap,
	}
}

func (s *IntStack) push(str int) {
	if s.size == s.cap {
		newData := make([]int, s.cap*2)
		copy(newData, s.data)
		s.data = newData
		s.cap *= 2
	}

	s.data[s.size] = str
	s.size++
}

func (s *IntStack) pop() int {
	s.size--
	return s.data[s.size]
}

func (s *IntStack) top() int {
	return s.data[s.size-1]
}

func TestIt(t *testing.T) {
	inputs := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{0},
	}
	results := [][]int{
		{1, 1, 4, 2, 1, 1, 0, 0},
		{0},
	}

	for i := 0; i < len(inputs); i++ {
		s := dailyTemperatures(inputs[i])
		result := results[i]
		for j := 0; j < len(s); j++ {
			if s[j] != result[j] {
				t.Errorf("expected %v, but %v", result, s)
			}
		}
	}
}
