package main

import (
	"testing"
)

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	return max(maxSumTwoNoOverlapAndOrdered(nums, firstLen, secondLen),
		maxSumTwoNoOverlapAndOrdered(nums, secondLen, firstLen))
}

func maxSumTwoNoOverlapAndOrdered(nums []int, firstLen int, secondLen int) int {
	lastI := len(nums) - secondLen
	firstSum, secondSum := 0, 0
	for i := 0; i < firstLen; i++ {
		firstSum += nums[i]
	}
	for i := firstLen; i < firstLen+secondLen; i++ {
		secondSum += nums[i]
	}

	maxFirstSum := firstSum
	maxSum := firstSum + secondSum
	for i := firstLen + 1; i <= lastI; i++ {
		firstSum = firstSum - nums[i-firstLen-1] + nums[i-1]
		secondSum = secondSum - nums[i-1] + nums[i+secondLen-1]
		maxFirstSum = max(maxFirstSum, firstSum)
		maxSum = max(maxSum, maxFirstSum+secondSum)
	}
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestIt(t *testing.T) {
	nums := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
	firstLen := 1
	secondLen := 2
	expect := 20

	result := maxSumTwoNoOverlap(nums, firstLen, secondLen)

	if result != expect {
		t.Errorf("expected %v, but %v", expect, result)
	}
}
