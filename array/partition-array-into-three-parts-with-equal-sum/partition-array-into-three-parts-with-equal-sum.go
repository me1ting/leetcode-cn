package main

// https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
// 两次扫描，先计算平均值，再检验
func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, n := range A {
		sum += n
	}

	if sum%3 != 0 {
		return false
	} else {
		subSum, last := sum/3, len(A)-1
		count, tmpSum := 0, 0
		for i, n := range A {
			tmpSum += n
			if tmpSum == subSum {
				tmpSum = 0
				count++
				if count == 2 && i < last {
					return true
				}
			}
		}
		return false
	}
}
