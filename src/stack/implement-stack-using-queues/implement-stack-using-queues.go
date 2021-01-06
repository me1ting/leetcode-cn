package implement_stack_using_queues

type MyStack struct {
	data  []int
	count int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{data: []int{}}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
	this.count++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	this.count--
	val := this.data[this.count]
	this.data = this.data[0:this.count]
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data[this.count-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.count == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
