package generateparentheses

import (
	"testing"
)

// https://leetcode.cn/problems/generate-parentheses
// 使用递归法解决该问题，存在两个规律：
// - 结果中最多只有n个左括号，n个右括号
// - 从左向右的扫描过程中，右括号的数量不能多余左括号的数量
func generateParenthesis(n int) []string {
	res := []string{}
	lCnt, rCnt := 0, 0
	heap := []byte{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n*2 {
			res = append(res, string(heap))
		}
		if lCnt < n {
			heap = append(heap, '(')
			lCnt++
			dfs(i + 1)
			heap = heap[:len(heap)-1]
			lCnt--
		}

		if rCnt < lCnt {
			heap = append(heap, ')')
			rCnt++
			dfs(i + 1)
			heap = heap[:len(heap)-1]
			rCnt--
		}
	}
	dfs(0)
	return res
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n = 1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n = 2",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "n = 3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n = 4",
			n:    4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := generateParenthesis(tt.n); !equalSets(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 从集合的角度比较两个 []string 是否相等
func equalSets(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	// 将第一个切片转换为集合
	setA := make(map[string]bool)
	for _, v := range a {
		setA[v] = true
	}

	// 检查第二个切片中的元素是否都在集合中
	for _, v := range b {
		if !setA[v] {
			return false
		}
	}

	return true
}

func TestEqualSets(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want bool
	}{
		{
			name: "相同元素，顺序不同",
			a:    []string{"a", "b", "c"},
			b:    []string{"c", "b", "a"},
			want: true,
		},
		{
			name: "相同元素，顺序相同",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "c"},
			want: true,
		},
		{
			name: "不同元素",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "d"},
			want: false,
		},
		{
			name: "长度不同",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b"},
			want: false,
		},
		{
			name: "空切片",
			a:    []string{},
			b:    []string{},
			want: true,
		},
		{
			name: "一个切片为空",
			a:    []string{"a", "b", "c"},
			b:    []string{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSets(tt.a, tt.b); got != tt.want {
				t.Errorf("equalSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
