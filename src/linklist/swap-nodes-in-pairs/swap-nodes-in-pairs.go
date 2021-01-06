package swap_nodes_in_pairs


type ListNode struct {
	Val  int
	Next *ListNode
}
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
// 没什么难度，基本情况就是保留前一个节点引用并进行交换，需要处理头和长度不够的额外情况
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := head
	if node.Next != nil {
		next := node.Next
		node.Next = next.Next
		next.Next = node
		head = next
	}

	prev := node
	node = prev.Next
	for node != nil && node.Next != nil {
		next := node.Next
		node.Next = next.Next
		next.Next = node
		prev.Next = next

		prev = node
		node = prev.Next
	}

	return head
}
