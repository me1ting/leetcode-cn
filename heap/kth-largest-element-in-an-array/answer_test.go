package answer

import (
	"container/heap"
	"testing"
)

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 堆的经典题目
// 做此题的目的是熟悉Go自带的heap库
func findKthLargest(nums []int, k int) int {
	var ih IntHeap = make([]int, k)
	copy(ih, nums[:k])
	h := &ih
	heap.Init(h)

	for i := k; i < len(nums); i++ {
		if nums[i] > ih[0] {
			ih[0] = nums[i]
			heap.Fix(h, 0)
		}
	}
	return ih[0]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestIt(t *testing.T) {
	result := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	if result != 5 {
		t.Errorf("error: expected %v, but %v", 5, result)
	}

	result = findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	if result != 4 {
		t.Errorf("error: expected %v, but %v", 4, result)
	}
}
