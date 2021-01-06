package answer

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
// 后序遍历：即按照左右根的顺序遍历树
// 参考中序遍历
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var traversal func(*TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Val)
	}

	traversal(root)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
