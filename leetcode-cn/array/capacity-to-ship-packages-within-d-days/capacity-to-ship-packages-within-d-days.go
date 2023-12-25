package main

import "fmt"

// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
// 一开始我想多了，什么方差、曲线之类的数学上的中位数什么的
// 唯一有效的信息是最大负载即船的最小负载，以及所有负载，在这范围内使用二分查找左闭区间即可
func shipWithinDays(weights []int, D int) int {
	max, sum := 0, 0
	for _, w := range weights {
		if w > max {
			max = w
		}
		sum += w
	}

	l, r := max, sum
	if match(weights, D, l) {
		return l
	}
	// 二分搜索左闭区间边界，l不匹配，r匹配
	for l+1 < r {
		center := (l + r) / 2
		if match(weights, D, center) {
			r = center
		} else {
			l = center
		}
	}
	return r
}

func match(weights []int, D int, capacity int) bool {
	d, load := 0, 0

	for _, w := range weights {
		if load += w; load > capacity {
			load = w
			d++

			if d == D {
				return false
			}
		}
	}

	return true
}

func main() {
	weights := []int{1,2,3,4,5,6,7,8,9,10}
	D := 5

	fmt.Println(match(weights,D,14))
}