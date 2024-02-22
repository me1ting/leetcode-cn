package constructbinarytreefrompreorderandpostordertraversal_test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	idx := map[int]int{}
	for i, n := range postorder {
		idx[n] = i
	}

	var build func(i, j, m, n int) *TreeNode

	build = func(i, j, m, n int) *TreeNode {
		if i == j {
			return nil
		}
		if i+1 == j {
			return &TreeNode{Val: preorder[i]}
		}

		val := preorder[i]
		c1, c2 := preorder[i+1], postorder[n-2]
		if c1 == c2 {
			return &TreeNode{
				Val:  val,
				Left: build(i+1, j, m, n-1),
			}
		}

		leftTreeSize := idx[c1] + 1 - m

		return &TreeNode{
			Val:   val,
			Left:  build(i+1, i+1+leftTreeSize, m, m+leftTreeSize),
			Right: build(i+1+leftTreeSize, j, m+leftTreeSize, n-1),
		}

		return nil
	}

	return build(0, len(preorder), 0, len(postorder))
}
