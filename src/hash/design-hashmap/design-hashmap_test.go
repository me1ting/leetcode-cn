package answer

import "testing"

// https://leetcode-cn.com/problems/design-hashmap/
// 测试始终不能通过，以后解决
type MyHashMap struct {
	size int
	cap  int
	data []*node
}

type node struct {
	next  *node
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		0,
		7777,
		make([]*node, 7777),
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	/*
		if this.size >= int(float32(this.cap)*0.75) {
			this.scale(this.cap * 2)
		}
	*/

	i := key % this.cap
	n := this.data[i]
	if n == nil {
		this.data[i] = &node{nil, key, value}
		this.size++
		return
	}
	for {
		if n.key == key {
			n.value = value
			return
		}
		if n.next == nil {
			break
		}
		n = n.next
	}

	n.next = &node{nil, key, value}
	this.size++
}

/*
func (this *MyHashMap) scale(newCap int) {
	newData := make([]*node, newCap)
	for _, curr := range this.data {
		for curr != nil {
			i := curr.key % newCap
			n := newData[i]
			if n == nil {
				newData[i] = &node{nil, curr.key, curr.value}
			} else {
				for n.next != nil {
					n = n.next
				}
				n.next = &node{nil, curr.key, curr.value}
			}
			curr = curr.next
		}
	}
	this.data = newData
	this.cap = newCap
}
*/

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	i := key % this.cap
	n := this.data[i]

	for n != nil {
		if n.key == key {
			return n.value
		}
		n = n.next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	/*
		if 4*this.size < this.cap && this.cap > 10 {
			newCap := this.cap / 2
			if newCap < 10 {
				newCap = 10
			}
			this.scale(newCap)
		}
	*/

	i := key % this.cap
	n := this.data[i]
	if n == nil {
		return
	}

	if n.key == key {
		this.data[i] = nil
		this.size--
		return
	}

	prev := n
	n = n.next
	for n != nil {
		if n.key == key {
			prev.next = n.next
			this.size--
			return
		}
		prev = n
		n = n.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func TestMyHashMap(t *testing.T) {
}
