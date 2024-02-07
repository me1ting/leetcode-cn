package removenodesfromlinkedlist_test

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	stack := Stack[*ListNode]([]*ListNode{})
	for node := head; node != nil; node = node.Next {
		for len(stack) > 0 && stack.top().Val < node.Val {
			stack.pop()
		}
		stack.push(node)
	}
	for i := 1; i < len(stack); i++ {
		stack[i-1].Next = stack[i]
	}
	return stack[0]
}

type Stack[T any] []T

func (s *Stack[T]) push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) pop() T {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s Stack[T]) top() T {
	return s[len(s)-1]
}
