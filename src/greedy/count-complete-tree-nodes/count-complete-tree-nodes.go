package count_complete_tree_nodes

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/count-complete-tree-nodes/
// 根据从数组堆里面学到的方法，我们可以很自然的想出：使用一个栈来记录路径，直到找到第一个空穴
// 根据数组堆索引的特点来计算空穴的索引，减1等于节点个数
// 进一步优化：
// 1.根据题目上下文（使用int表示节点个数），我们可以得知level最大为31，可以使用一个带符号整数来表示路径
// 2.不需要依次查找所有可能的空穴位置，使用二分查找进行加速，二分查找需要O((log_2(n))^2)时间，而遍历需要O(N)
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	for node := root; node != nil; node = node.Left {
		level++
	}

	full := math.MaxInt32
	if level < 31 {
		full = 1<<level - 1
	}
	if count := countNodesInner(root, 1<<(level-1), 1); count > 0 {
		return count
	}

	return full
}

// start：最底层的开始节点索引
func countNodesInner(root *TreeNode, start int, index int) int {
	if index >= start {
		if root == nil {
			return index - 1
		} else {
			return -1
		}
	}
	if n := countNodesInner(root.Left, start, index<<1); n > 0 {
		return n
	}
	return countNodesInner(root.Right, start, index<<1+1)
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	for node := root; node != nil; node = node.Left {
		level++
	}

	start := 1 << (level - 1)
	end := 1<<level
	//避免溢出
	if level == 31 {
		node:=root
		for i:=1;i<32;i++{
			node = node.Right
		}
		if node !=nil {
			return math.MaxInt32
		}
		end = math.MaxInt32
	}

	return sort.Search(end, func(i int) bool {
		if i <= start {
			return false
		}
		// 如何根据索引重建数组堆中的路径是很重要的技巧
		// 第n层引入的+1（右分支）并不会受到后续的走向的影响，可以通过位运算提取出来
		// (((1<<1 + ?)<<1 + ?)...)<<1 + ?
		// 因为移位的原因，每次在末尾的0,1选择并不会受到任何影响
		bits := 1 << (level-2)
		node := root
		for node != nil && bits > 0 {
			if bits&i == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}

		return node == nil
	}) - 1
}
