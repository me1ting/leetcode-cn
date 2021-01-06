package last_stone_weight

import (
	"testing"
)

// https://leetcode-cn.com/problems/last-stone-weight/
// 使用堆来解决问题
// 从算法效率角度来讲，最佳的方式是在数组上直接构建堆
// 但是从工程角度来讲，使用堆API是更合理的方式
// 这里选择后者
func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap(len(stones))
	for _, w := range stones {
		heap.insert(w)
	}

	for heap.size > 1 {
		w1 := heap.removeMax()
		w2 := heap.removeMax()
		if w1 > w2 {
			heap.insert(w1 - w2)
		}
	}

	if heap.size == 1 {
		return heap.removeMax()
	}

	return 0
}

type MaxHeap struct {
	data      []int
	cap, size int
}

func NewMaxHeap(cap int) *MaxHeap {
	return &MaxHeap{
		data: make([]int, cap+1, cap+1),
		cap:  cap,
		size: 0,
	}
}

func (m *MaxHeap) insert(val int) {
	m.size++
	m.data[m.size] = val
	m.swim(m.size)
}

func (m *MaxHeap) removeMax() int {
	val := m.data[1]
	m.data[0] = m.data[m.size]
	m.size--
	m.sink(0)
	return val
}

func (m *MaxHeap) swim(i int) {
	data := m.data[i]
	for i > 1 && m.data[i/2] < data {
		m.data[i] = m.data[i/2]
		i = i / 2
	}
	m.data[i] = data
}

func (m *MaxHeap) sink(i int) {
	data := m.data[i]
	for 2*i <= m.size {
		j := 2 * i
		if j < m.size && m.data[j] < m.data[j+1] {
			j++
		}
		if m.data[j] <= data {
			break
		}
		m.data[i] = m.data[j]
		i = j
	}
	m.data[i] = data
}

func TestIt(t *testing.T) {
	testcase := []int{2, 7, 4, 1, 8, 1}
	if lastStoneWeight(testcase) != 1 {
		t.Errorf("error")
	}
}
