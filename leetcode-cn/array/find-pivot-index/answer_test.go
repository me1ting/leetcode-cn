package answer

// https://leetcode-cn.com/problems/find-pivot-index/
// 一道简单题，但是很难独立第一时间想到如何找到中心索引
// 方法其实很简单，先统计所有元素的总和，找到sum_curr*2+num_curr==sum的点。
func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	sumCurr := 0
	for i, num := range nums {
		if sumCurr*2+num == sum {
			return i
		}
		sumCurr += num
	}
	return -1
}
