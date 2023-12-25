package linked_list_cycle_ii

import (
	"fmt"
	"testing"
)

// 问题：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/
// 题解：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
// 快慢指针可以找到环，但是判断入口需要其它方法。
// 而这方法来源于数学计算，参考题解。

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slow, fast = head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			break
		}
	}

	fast = head
	for fast != slow {
		slow, fast = slow.Next, fast.Next
	}
	return fast
}

func TestDetect(t *testing.T) {
	head := &ListNode{
		Val:  3,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next.Next = &ListNode{
		Val:  0,
		Next: nil,
	}
	head.Next.Next.Next = &ListNode{
		Val:  -4,
		Next: head.Next,
	}

	fmt.Println(detectCycle(head).Val)
}
