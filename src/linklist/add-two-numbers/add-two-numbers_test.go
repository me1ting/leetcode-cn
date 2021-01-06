package add_two_numbers

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/add-two-numbers/submissions/
// 没什么难度
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry = 0
	var m, n *ListNode

	for m, n = l1, l2; ; m, n = m.Next, n.Next {
		m.Val += n.Val + carry
		if m.Val >= 10 {
			m.Val %= 10
			carry = 1
		} else {
			carry = 0
		}
		if m.Next == nil || n.Next == nil {
			break
		}
	}

	if n.Next != nil {
		m.Next = n.Next
	}

	if m.Next != nil {
		for m = m.Next; ; m = m.Next {
			m.Val += carry
			if m.Val >= 10 {
				m.Val %= 10
				carry = 1
			} else {
				carry = 0
			}
			if m.Next == nil {
				break
			}
		}
	}

	if carry > 0 {
		m.Next = &ListNode{
			Val:  carry,
			Next: nil,
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

