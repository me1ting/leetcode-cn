package add_two_numbers_backup

import "testing"

// 这是一份针对顺序链表实现的代码，可能后续会遇到该问题，因此保留
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersInner(l1 *ListNode, l2 *ListNode, pos, len1, len2 int) int {
	if pos == 0 {
		return 0
	}

	if pos > len2 {
		carry := addTwoNumbersInner(l1.Next, l2, pos-1, len1, len2)
		l1.Val += carry
	} else {
		carry := addTwoNumbersInner(l1.Next, l2.Next, pos-1, len1, len2)
		l1.Val += l2.Val + carry
	}

	if l1.Val >= 10 {
		l1.Val %= 10
		return 1
	}

	return 0
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var len1, len2 int
	for i := l1; i != nil; i = i.Next {
		len1++
	}
	for i := l2; i != nil; i = i.Next {
		len2++
	}

	if len1 < len2 {
		l1, l2 = l2, l1
		len1, len2 = len2, len1
	}

	if carry := addTwoNumbersInner(l1, l2, len1, len1, len2); carry > 0 {
		return &ListNode{
			Val:  carry,
			Next: l1,
		}
	}

	return l1
}

func TestAdd(t *testing.T) {
	l1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	addTwoNumbers(l1,l2)
}
