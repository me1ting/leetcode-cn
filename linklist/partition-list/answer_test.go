package answer

// https://leetcode-cn.com/problems/partition-list/
// 根据直觉可以很快得出答案
func partition(head *ListNode, x int) *ListNode {
	smallHead, bigHead := &ListNode{}, &ListNode{}
	smallTail, bigTial := smallHead, bigHead

	for node := head; node != nil; node = node.Next {
		if node.Val <= x {
			smallTail.Next = node
			smallTail = node
		} else {
			bigTial.Next = node
			bigTial = node
		}
	}

	smallTail.Next = bigHead.Next
	bigTial.Next = nil

	return smallHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
