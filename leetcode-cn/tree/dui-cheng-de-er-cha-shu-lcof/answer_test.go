package answer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/
// 和https://leetcode-cn.com/problems/same-tree/是等价题，使用DFS即可
// 转换为判断两个子树是否是镜像
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isReverse func(*TreeNode, *TreeNode) bool
	isReverse = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}

		if t1.Val != t2.Val {
			return false
		}
		return isReverse(t1.Left, t2.Right) && isReverse(t1.Right, t2.Left)
	}

	return isReverse(root.Left, root.Right)
}
