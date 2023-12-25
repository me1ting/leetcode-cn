package answer

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// 中序遍历，指root位于中间位置
// 最直观的方式是使用递归
// 也可以使用迭代，但代码难以理解，个人看来也不用理解
// 迭代示例代码：
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio/
func inorderTraversal(root *TreeNode) (res []int) {
	var traversal func(*TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}

	traversal(root)
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
