package answer

import (
	"container/heap"
	"testing"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// min-heap
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
// 流的第k大数，最佳方式是使用堆实现
type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	var h IntHeap = make([]int, 0)
	heap.Init(&h)
	for i := 0; i < k && i < len(nums); i++ {
		heap.Push(&h, nums[i])
	}

	for i := k; i < len(nums); i++ {
		heap.Push(&h, nums[i])
		heap.Pop(&h)
	}

	return KthLargest{
		heap: &h,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}

func TestKthLargest(t *testing.T) {
	kl := Constructor(3, []int{4, 5, 8, 2})
	testcases := []int{3, 5, 10, 9, 4}
	expecteds := []int{4, 5, 5, 8, 8}
	for i, testcase := range testcases {
		expected := expecteds[i]
		val := kl.Add(testcase)
		if val != expected {
			t.Errorf("return %d, expected %d", val, expected)
			break
		}
	}

	kl = Constructor(1, []int{})
	kl.Add(-3)
}
