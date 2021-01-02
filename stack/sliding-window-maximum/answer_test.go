package answer

import "testing"

// https://leetcode-cn.com/problems/sliding-window-maximum/
// 题解 https://leetcode-cn.com/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
// 直接思维：记录窗口最大值，当滑动窗口移动时判断是否移除，但是次大值呢？。。。
// 1. 优先队列+索引（使用索引间接读取值是这个算法最值得学习的技巧）：
//   使用优先队列存储窗口，但是存储的是索引，从而间接读取值，这样可以在窗口移动时忽略失效的值（根据索引）
//   每次移动窗口，入堆最新值的索引，并弹出失效的最大值，最后记录最大值
//
// 2. 单调队列+索引：
// 这里单调队列指双端队列存储着单调递增/递减的值
// 使用单调队列记录窗口，但存储的是索引，从而间接读取值
// 每次移动窗口，入队最新值的索引，并直接丢弃更小值的索引（它们作为过去的小值对于最大值没有影响）
// ，出队失效的最大值，最后记录最大值
//
// 这里使用单调队列来实现
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	n := len(nums)
	ans := make([]int, 0, n-k+1)
	q := NewRingQueue(k)

	pushR := func(i int) {
		for q.Size() > 0 && nums[i] >= nums[q.RFetch()] {
			q.RPop()
		}
		q.RPush(i)
	}

	popL := func(i int) {
		for q.LFetch() <= i-k {
			q.LPop()
		}
	}

	for i := 0; i < k; i++ {
		pushR(i)
	}
	ans = append(ans, nums[q.LFetch()])

	for i := k; i < n; i++ {
		popL(i)
		pushR(i)
		ans = append(ans, nums[q.LFetch()])
	}

	return ans
}

// 这里使用环来实现一个固定容量的双端队列
type RingQueue struct {
	data []int
	cap  int
	size int
	l, r int //有效元素的第一个位置，下一个存储位置
}

// 为了简便，这里不考虑参数有效性
func NewRingQueue(cap int) *RingQueue {
	return &RingQueue{
		make([]int, cap),
		cap,
		0,
		0,
		0,
	}
}

// 为了简便，这里不考虑边界情况：队列已满
func (q *RingQueue) RPush(i int) {
	q.data[q.r] = i
	q.r++
	if q.r == q.cap {
		q.r = 0
	}
	q.size++
}

// 为了简便，这里不考虑边界情况：队列为空
func (q *RingQueue) RPop() int {
	q.r--
	if q.r == -1 {
		q.r += q.cap
	}
	q.size--
	return q.data[q.r]
}

// 为了简便，这里不考虑边界情况：队列为空
func (q *RingQueue) RFetch() int {
	i := q.r - 1
	if i == -1 {
		i += q.cap
	}
	return q.data[i]
}

// 为了简便，这里不考虑边界情况：队列为空
func (q *RingQueue) LPop() int {
	e := q.data[q.l]
	q.l++
	if q.l == q.cap {
		q.l = 0
	}
	q.size--
	return e
}

// 为了简便，这里不考虑边界情况：队列为空
func (q *RingQueue) LFetch() int {
	return q.data[q.l]
}

func (q *RingQueue) Size() int {
	return q.size
}

func TestIt(t *testing.T) {
	inputs1 := [][]int{
		{1, 3, -1, -3, 5, 3, 6, 7},
		{1, -1},
	}
	inputs2 := []int{
		3,
		1,
	}
	expecteds := [][]int{
		{3, 3, 5, 5, 6, 7},
		{1, -1},
	}

	for i := 0; i < len(inputs1); i++ {
		nums := inputs1[i]
		k := inputs2[i]
		expected := expecteds[i]

		result := maxSlidingWindow(nums, k)

		if len(result) != len(expected) {
			t.Errorf("error")
			return
		}

		for i := 0; i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Errorf("error")
				return
			}
		}
	}
}
