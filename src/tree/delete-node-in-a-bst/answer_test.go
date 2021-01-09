package main

// https://leetcode-cn.com/problems/delete-node-in-a-bst
// 二叉搜索树的基本操作
// 删除一个节点：
// 1.如果是叶节点，直接删除
// 2.使用左子树的最大值（前驱节点）或者右子树的最小值（后继节点）替代，并递归删除它，始终选择一边会导致平衡性问题，但这里不用考虑过多
//   由于查找前驱节点或者后继节点，然后递归删除存在重复操作，可以合二为一
func deleteNode(root *TreeNode, key int) *TreeNode {
	var parent *TreeNode = nil
	node := root

	for node != nil && node.Val != key {
		parent = node
		if key < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if node == nil {
		return root
	}

	if node.Left != nil {
		node.Val = deletePredecessor(node)
	} else if node.Right != nil {
		node.Val = deleteSuccessor(node)
	} else {
		if parent == nil {
			return nil
		}

		if node == parent.Left {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}

	return root
}

// 要求左子树存在
func deletePredecessor(root *TreeNode) int {
	node := root.Left
	if node.Right != nil {
		parent := root
		for node.Right != nil {
			parent = node
			node = node.Right
		}
		parent.Right = node.Left
		node.Left = nil
	} else {
		root.Left = node.Left
		node.Left = nil
	}

	return node.Val
}

// 要求右子树存在
func deleteSuccessor(root *TreeNode) int {
	node := root.Right
	if node.Left != nil {
		parent := root
		for node.Left != nil {
			parent = node
			node = node.Left
		}
		parent.Left = node.Right
		node.Right = nil
	} else {
		root.Right = node.Right
		node.Right = nil
	}

	return node.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
