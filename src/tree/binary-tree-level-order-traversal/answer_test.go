package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	var visit func(*TreeNode, int)
	visit = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(ans) < level+1 {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], node.Val)
		visit(node.Left, level+1)
		visit(node.Right, level+1)
	}

	visit(root, 0)

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
