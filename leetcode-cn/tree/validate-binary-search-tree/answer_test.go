package validatebinarysearchtree_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/validate-binary-search-tree/
// 思路：
// 由于没有参考题解，所以自己的解题代码显得有些繁琐，但基本思路与题解的中的递归解法一致
// 由于题目限制val的有效范围为int32，题解使用的是int类型在Go x64上是int64，因此不存在问题
func isValidBST(root *TreeNode) bool {
	var smallerThan func(*TreeNode, int) bool
	var biggerThan func(*TreeNode, int) bool
	var between func(*TreeNode, int, int) bool

	smallerThan = func(root *TreeNode, v int) bool {
		if root == nil {
			return true
		}
		if root.Val >= v {
			return false
		}

		return smallerThan(root.Left, root.Val) && between(root.Right, root.Val, v)
	}

	biggerThan = func(root *TreeNode, v int) bool {
		if root == nil {
			return true
		}
		if root.Val <= v {
			return false
		}

		return between(root.Left, v, root.Val) && biggerThan(root.Right, root.Val)
	}

	between = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val <= min || root.Val >= max {
			return false
		}

		return between(root.Left, min, root.Val) && between(root.Right, root.Val, max)
	}

	return smallerThan(root.Left, root.Val) && biggerThan(root.Right, root.Val)
}
