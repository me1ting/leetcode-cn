package answer

import (
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/simplify-path/
func simplifyPath(path string) string {
	tokens := strings.Split(path, "/")
	stack := NewStack(10)

	for _, token := range tokens {
		switch token {
		case "":
		case ".":
		case "..":
			if stack.size > 0 {
				stack.pop()
			}
		default:
			stack.push(token)
		}
	}

	return "/" + strings.Join(stack.data[:stack.size], "/")
}

type Stack struct {
	data []string
	size int
	cap  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]string, cap),
		size: 0,
		cap:  cap,
	}
}

func (s *Stack) push(str string) {
	if s.size == s.cap {
		newData := make([]string, s.cap*2)
		copy(newData, s.data)
		s.data = newData
		s.cap *= 2
	}

	s.data[s.size] = str
	s.size++
}

func (s *Stack) pop() string {
	s.size--
	return s.data[s.size]
}

func (s *Stack) top() string {
	return s.data[s.size-1]
}

func TestIt(t *testing.T) {
	inputs := []string{
		"/home/",
		"/../",
		"/home//foo/",
		"/a/./b/../../c/",
		"/a/../../b/../c//.//",
		"/a//b////c/d//././/..",
	}
	results := []string{
		"/home",
		"/",
		"/home/foo",
		"/c",
		"/c",
		"/a/b/c",
	}

	for i := 0; i < len(inputs); i++ {
		s := simplifyPath(inputs[i])
		if s != results[i] {
			t.Errorf("expected %s, but %s", results[i], s)
		}
	}
}
