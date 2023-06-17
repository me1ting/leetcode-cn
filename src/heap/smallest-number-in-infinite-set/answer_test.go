package answer

import "container/heap"

type SmallestInfiniteSet struct {
	minOfSlice int
	heap       IntHeap
	numInHeap  map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{1, []int{}, map[int]bool{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.heap.Len() == 0 {
		num := this.minOfSlice
		this.minOfSlice++
		return num
	}

	num := heap.Pop(&this.heap).(int)
	delete(this.numInHeap, num)
	return num
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.minOfSlice {
		return
	}

	if this.numInHeap[num] {
		return
	}

	if num == this.minOfSlice-1 {
		this.minOfSlice--
		return
	}

	heap.Push(&this.heap, num)
	this.numInHeap[num] = true
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
