package minimumoperationstoexceedthresholdvalueii

import (
	"container/heap"
	"testing"
)

func minOperations(nums []int, k int) int {
	intHeap := IntHeap(nums)
	h := &intHeap
	heap.Init(h)

	cnt := 0
	for (*h)[0] < k {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		heap.Push(h, x<<1+y)
		cnt++
	}

	return cnt
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestMinOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 11, 10, 1, 3}, 10, 2},
		{[]int{1, 1, 2, 4, 9}, 20, 4},
	}

	for _, tt := range tests {
		got := minOperations(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("minOperations(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
		}
	}
}
