package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// 根据前序遍历和中序遍历的特点，使用递归重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	inorderIndex := findOrder(inorder,rootVal)
	return &TreeNode{
		Val:rootVal,
		Left: buildTree(preorder[1:1+inorderIndex],inorder[:inorderIndex]),
		Right: buildTree(preorder[1+inorderIndex:],inorder[1+inorderIndex:]),
	}
}

func findOrder(inorder []int, val int) int {
	for i,v:=range inorder {
		if v == val {
			return i
		}
	}
	return -1
}
