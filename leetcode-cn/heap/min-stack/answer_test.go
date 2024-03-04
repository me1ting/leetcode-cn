package minstack_test

// 简单题
type MinStack struct {
	data     []int
	min, len int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.len++

	if this.len == 1 {
		this.min = val
	} else if this.min > val {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	this.len--
	len := this.len
	val := this.data[len]
	this.data = this.data[:len]

	min := this.min
	if len != 0 && val == min {
		min = this.data[0]
		for _, n := range this.data {
			if n < min {
				min = n
			}
		}
		this.min = min
	}
}

func (this *MinStack) Top() int {
	return this.data[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
