package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表
// 使用`快慢指针`经典方法
// https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var fast, slow = head.Next, head
	var slowTime = false
	for fast != nil && fast != slow {
		if slowTime {
			slow = slow.Next
			slowTime = false
		} else {
			slowTime = true
		}
		fast = fast.Next

	}
	return fast == slow
}
