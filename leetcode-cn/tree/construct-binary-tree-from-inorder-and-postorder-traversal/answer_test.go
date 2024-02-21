package constructbinarytreefrominorderandpostordertraversal_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// 本题是 construct-binary-tree-from-preorder-and-inorder-traversal 的镜像版本
// 后序遍历：先遍历左子树，再遍历右子树，最后打印节点
func buildTree(inorder []int, postorder []int) *TreeNode {
	idx := map[int]int{}
	for i, n := range inorder {
		idx[n] = i
	}

	var build func(int, int, int, int) *TreeNode
	build = func(startPost, endPost, startIn, endIn int) *TreeNode {
		if startIn == endIn {
			return nil
		}

		val := postorder[endPost-1]
		pos := idx[val]
		leftTreeSize := pos - startIn

		return &TreeNode{
			Val:   val,
			Left:  build(startPost, startPost+leftTreeSize, startIn, pos),
			Right: build(startPost+leftTreeSize, endPost-1, pos+1, endIn),
		}
	}

	return build(0, len(postorder), 0, len(inorder))
}
