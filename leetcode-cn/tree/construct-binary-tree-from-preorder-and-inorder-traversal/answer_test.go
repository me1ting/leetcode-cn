package constructbinarytreefrompreorderandinordertraversal_test

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
// 先序遍历：先输出节点值，再遍历左子树，再遍历右子树
// 中序遍历：先遍历左子树，再输出节点值，最后遍历右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	idx := map[int]int{}
	for i, n := range inorder {
		idx[n] = i
	}

	var build func(int, int, int, int) *TreeNode
	build = func(startPre, endPre, startIn, endIn int) *TreeNode {
		if startIn == endIn {
			return nil
		}

		val := preorder[startPre]
		pos := idx[val]
		leftTreeSize := pos - startIn

		return &TreeNode{
			Val:   val,
			Left:  build(startPre+1, startPre+1+leftTreeSize, startIn, pos),
			Right: build(startPre+1+leftTreeSize, endPre, pos+1, endIn),
		}
	}

	return build(0, len(preorder), 0, len(inorder))
}

func TestIt(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	buildTree(preorder, inorder)
}
