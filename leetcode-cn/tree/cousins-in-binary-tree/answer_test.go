package cousinsinbinarytree_test

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
func isCousins(root *TreeNode, x int, y int) bool {
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		children := []*TreeNode{}
		xExist, yExist := false, false // 注意，x,y可能同值
		for _, node := range nodes {
			newXExist, newYExist := false, false
			if node.Left != nil {
				val := node.Left.Val
				if val == x && yExist || val == y && xExist {
					return true
				}
				if val == x {
					newXExist = true
				} else if val == y {
					newYExist = true
				}
				children = append(children, node.Left)
			}
			if node.Right != nil {
				val := node.Right.Val
				if val == x && yExist || val == y && xExist {
					return true
				}
				if val == x {
					newXExist = true
				} else if val == y {
					newYExist = true
				}
				children = append(children, node.Right)
			}

			if newXExist {
				xExist = true
			}
			if newYExist {
				yExist = true
			}
		}
		nodes = children
	}

	return false
}