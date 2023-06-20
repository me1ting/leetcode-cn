package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	stack := []*ListNode{}

	for node := head; node != nil; node = node.Next {
		stack = append(stack, node)
	}

	for n := len(stack) - 1; n > 0; n-- {
		stack[n].Next = stack[n-1]
	}

	stack[0].Next = nil

	return stack[len(stack)-1]
}

func TestReverseList(t *testing.T) {
	virHead := &ListNode{
		0,
		nil,
	}
	prev := virHead
	for i := 1; i <= 5; i++ {
		prev.Next = &ListNode{
			i,
			nil,
		}
		prev = prev.Next
	}

	head := reverseList(virHead.Next)

	stack := []int{}
	for node := head; node != nil; node = node.Next {
		stack = append(stack, node.Val)
	}

	for i, n := range stack {
		if i+n != 5 {
			t.Error("bad ans")
		}
	}
}
