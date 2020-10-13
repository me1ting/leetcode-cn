package happy_number

import "testing"

// https://leetcode-cn.com/problems/happy-number
// 官方题解：https://leetcode-cn.com/problems/happy-number/solution/kuai-le-shu-by-leetcode-solution/
// 此题的关键在于结果只存在两种情况：
//   1. 收敛为1
//   2. 在3位数里面循环
// 至于这个结论是怎么得出的，我也不知道
// 事实上只存在唯一一个循环，甚至可以硬编码这个循环来加快运算速度
func isHappy(n int) bool {
	record:=map[int]bool{}
	for n!=1 {
		if n<999 {
			if _,ok:=record[n];ok {
				return false
			}
		}
		record[n] = true
		n = next(n)
	}
	return true
}

func next(n int) int {
	sum:=0
	for n>0 {
		j:=n%10
		sum+=j*j
		n/=10
	}
	return sum
}

func TestIsHappy(t *testing.T) {
	n:=19
	if !isHappy(n) {
		t.Errorf("expected true, but false")
	}
}
