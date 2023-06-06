package answer

import "math/rand"

type RandomizedSet struct {
	data  []int
	table map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data:  make([]int, 0, 10),
		table: map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.table[val]; !ok {
		this.data = append(this.data, val)
		this.table[val] = len(this.data) - 1
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.table[val]; ok {
		size := len(this.data)
		if pos != size-1 {
			lastVal := this.data[size-1]
			this.data[pos] = lastVal
			this.table[lastVal] = pos
		}
		this.data = this.data[:size-1]
		delete(this.table, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	pos := rand.Intn(len(this.data))
	return this.data[pos]
}
