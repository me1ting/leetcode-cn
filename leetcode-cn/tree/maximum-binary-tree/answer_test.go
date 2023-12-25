package answer

import (
	"sort"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/maximum-binary-tree/
// 官方题解是使用递归
// 个人比较喜欢使用人类的思维方式来解决这个问题：
// 先排序，然后从大到小递归排序后的数组，根据索引位置来生长树
// 时间复杂度为O(n*2)
//
// 执行用时：40 ms, 在所有 Go 提交中击败了6.07%的用户
// 内存消耗：7.1 MB, 在所有 Go 提交中击败了5.14%的用
func constructMaximumBinaryTree1(nums []int) *TreeNode {
	indexs := map[int]int{}
	for i, n := range nums {
		indexs[n] = i
	}

	sort.Sort(sort.IntSlice(nums))
	root := &TreeNode{Val: nums[len(nums)-1]}
	for i := len(nums) - 2; i >= 0; i-- {
		node := root
		n := nums[i]
		index := indexs[n]

		for {
			if index < indexs[node.Val] {
				if node.Left == nil {
					node.Left = &TreeNode{
						Val: n,
					}
					break
				}
				node = node.Left
			} else {
				if node.Right == nil {
					node.Right = &TreeNode{
						Val: n,
					}
					break
				}
				node = node.Right
			}
		}
	}
	return root
}

// 官方题解
// 执行用时：20 ms, 在所有 Go 提交中击败了67.41%的用户
// 内存消耗：6.7 MB, 在所有 Go 提交中击败了96.46%的用户
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums)
}

func construct(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := maxIndex(nums)
	tree := &TreeNode{Val: nums[i]}
	tree.Left = construct(nums[:i])
	tree.Right = construct(nums[i+1:])

	return tree
}

func maxIndex(nums []int) int {
	max := 0
	m := nums[0]
	for i, n := range nums {
		if m < n {
			max = i
			m = n
		}
	}
	return max
}

func TestIt(t *testing.T) {
	testcase := []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(testcase)
}
