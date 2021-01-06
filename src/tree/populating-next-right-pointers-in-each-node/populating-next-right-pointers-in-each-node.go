package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
// 使用一个数组/stack...记录每一层的prev node，遍历时设置prev.Next = curr。
func connect(root *Node) *Node {
	level := 1
	preOrder(root, &[]*Node{nil}, level)
	return root
}

func preOrder(root *Node, prev *[]*Node, level int) {
	if root == nil {
		return
	}
	if len(*prev) <= level {
		*prev = append(*prev, root)
	} else {
		p := (*prev)[level]
		p.Next = root
		(*prev)[level] = root
	}
	preOrder(root.Left, prev, level+1)
	preOrder(root.Right, prev, level+1)
}
