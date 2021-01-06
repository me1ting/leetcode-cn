package main

// https://leetcode-cn.com/problems/teemo-attacking/
//　题目很简单，这里说说两者思维方式

/*
// 直接模拟，累积计时
func findPoisonedDuration(timeSeries []int, duration int) int {
	sum := 0
	start := 0
	end := 0
	for _,t := range timeSeries {
		if end < t {
			sum += end-start
			start = t
		}
		end = t + duration
	}

	return sum+end-start
}
*/

// 通过比较持续时间和间隔时间，取最小值累加，这经过了思维转换
// “好处”是性能略有提高
func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	sum := 0
	for i := 1; i < len(timeSeries); i++ {
		sum += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	return sum + duration
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
