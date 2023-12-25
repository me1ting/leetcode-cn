package miceandcheese

import (
	"container/heap"
	"testing"
)

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		diff[i] = reward1[i] - reward2[i]
	}

	if k == 0 {
		return sum(reward2)
	}

	if k == n {
		return sum(reward1)
	}

	h := IntHeap(diff[:k])
	heap.Init(&h)
	for i := k; i < n; i++ {
		d := diff[i]
		if d > h[0] {
			heap.Pop(&h)
			heap.Push(&h, d)
		}
	}

	return sum(reward2) + sum(h)
}

func sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	pos := len(*h) - 1
	val := (*h)[pos]
	*h = (*h)[:pos]
	return val
}

func TestMiceAndCheese(t *testing.T) {
	reward1List := [][]int{
		{1, 1, 3, 4},
	}
	reward2List := [][]int{
		{4, 4, 1, 1},
	}
	kList := []int{2}
	rightAnsList := []int{15}

	for i := 0; i < len(kList); i++ {
		reward1 := reward1List[i]
		reward2 := reward2List[i]
		k := kList[i]
		rightAns := rightAnsList[i]

		ans := miceAndCheese(reward1, reward2, k)
		if ans != rightAns {
			t.Errorf("expected %v, but got %v", rightAns, ans)
		}
	}
}
