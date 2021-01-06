package trapping_rain_water_test

import (
	"fmt"
	"testing"
)

// 求陷阱容量
// 求解这类问题最大的问题在于找到一个解题的模式，而非某些零散的思考点
// 单调栈就是求解该问题的一个模式：
// https://leetcode-cn.com/problems/trapping-rain-water/solution/dan-diao-zhan-jie-jue-jie-yu-shui-wen-ti-by-sweeti/
func trap(height []int) int {
	var sum = 0
	var stack = make([]int, len(height))
	// stack pointer
	var sp = -1
	for i, h := range height {
		if sp < 0 {
			sp++
			stack[sp] = i
		} else {
			if h < height[stack[sp]] {
				sp++
				stack[sp] = i
			} else if h == height[stack[sp]] {
				stack[sp] = i
			} else {
				n := stack[sp]
				for sp > 0 && h >= height[n] {
					m := stack[sp-1]
					if h > height[n] {
						fmt.Println(m, n, i)
						var diff int
						if h > height[m] {
							diff = height[m] - height[n]
						} else {
							diff = h - height[n]
						}
						sum += diff * (i - (m + 1))
					}
					n = m
					sp--
				}
				if sp == 0 && h >= height[n] {
					stack[sp] = i
				} else {
					sp++
					stack[sp] = i
				}
			}
		}
	}

	return sum
}

func TestTrap(t *testing.T) {
	testcase := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//testcase := []int{4, 2, 3}
	fmt.Println(trap(testcase))
}
