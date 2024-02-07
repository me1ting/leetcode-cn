package cousinsinbinarytreeii_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	sumMap := map[int]int{}

	var visit1 func(node *TreeNode, depth int)

	visit1 = func(node *TreeNode, depth int) {
		if node.Left != nil {
			visit1(node.Left, depth+1)
		}
		if node.Right != nil {
			visit1(node.Right, depth+1)
		}
		sumMap[depth] += node.Val
	}

	var visit2 func(node *TreeNode, depth int)
	visit2 = func(node *TreeNode, depth int) {
		sumChildren := 0
		if node.Left != nil {
			visit2(node.Left, depth+1)
			sumChildren += node.Left.Val
		}
		if node.Right != nil {
			visit2(node.Right, depth+1)
			sumChildren += node.Right.Val
		}

		childrenVal := sumMap[depth+1] - sumChildren
		if node.Left != nil {
			node.Left.Val = childrenVal
		}
		if node.Right != nil {
			node.Right.Val = childrenVal
		}
	}

	visit1(root, 1)
	visit2(root, 1)

	root.Val = 0

	return root
}
