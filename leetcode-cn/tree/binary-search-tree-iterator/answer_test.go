package answer

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-search-tree-iterator/
// 参考题解：https://leetcode-cn.com/problems/binary-search-tree-iterator/solution/er-cha-sou-suo-shu-die-dai-qi-by-leetcod-4y0y/
//
// 使用栈来遍历二叉树，但要实现足够精简的代码，需要一些技巧（套路，样板代码）
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (it *BSTIterator) Next() int {
	for node := it.cur; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}
	it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
	val := it.cur.Val
	it.cur = it.cur.Right
	return val
}

func (it *BSTIterator) HasNext() bool {
	return it.cur != nil || len(it.stack) > 0
}

/*
type BSTIterator struct {
	stack       []*TreeNode
	visited     []int
	next        *TreeNode
	root        *TreeNode
	rootVisited bool
}

func Constructor(root *TreeNode) BSTIterator {
	stack := []*TreeNode{}
	visited := []int{}

	return BSTIterator{
		stack,
		visited,
		nil,
		root,
		false,
	}
}

func (this *BSTIterator) nextInner() *TreeNode {
	if !this.rootVisited && this.root != nil {
		this.rootVisited = true
		this.stack = append(this.stack, this.root)
		this.visited = append(this.visited, 0)
	}

	i := len(this.stack) - 1
	if i >= 0 {
		curr := this.stack[i]
		visited := this.visited[i]
		if visited == 0 {
			this.visited[i] = 1
			left := curr.Left
			if left != nil {
				this.stack = append(this.stack, left)
				this.visited = append(this.visited, 0)
			}
			return this.nextInner()
		}

		if visited == 1 {
			this.visited[i] = 2
			right := curr.Right
			if right != nil {
				this.stack = append(this.stack, right)
				this.visited = append(this.visited, 0)
			}

			return curr
		}

		if visited == 2 {
			this.stack = this.stack[:i]
			this.visited = this.visited[:i]
			return this.nextInner()
		}
	}

	return nil
}

func (this *BSTIterator) Next() int {
	if this.next != nil {
		val := this.next.Val
		this.next = nil
		return val
	}

	next := this.nextInner()
	return next.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.next != nil {
		return true
	}

	node := this.nextInner()
	this.next = node

	return node != nil
}
*/

func TestIterator(t *testing.T) {
	seven := &TreeNode{
		7,
		nil,
		nil,
	}
	three := &TreeNode{
		3,
		nil,
		nil,
	}
	fifteen := &TreeNode{
		15,
		nil,
		nil,
	}
	nine := &TreeNode{
		9,
		nil,
		nil,
	}
	twenty := &TreeNode{
		20,
		nil,
		nil,
	}

	fifteen.Left = nine
	fifteen.Right = twenty
	seven.Left = three
	seven.Right = fifteen

	iter := Constructor(seven)

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
