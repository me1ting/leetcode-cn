package longest_mountain_in_array

import (
	"log"
	"testing"
)

// https://leetcode-cn.com/problems/longest-mountain-in-array/submissions/
// 一次遍历即可
func longestMountain(A []int) int {
	start := 0
	onMountain, up := false, true
	maxLen := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] == A[i+1] {
			if onMountain && !up {
				maxLen = max(maxLen, i-start+1)
			}
			onMountain = false
		} else if A[i] < A[i+1] {
			if !onMountain {
				start = i
				onMountain, up = true, true
			} else {
				if !up {
					maxLen = max(maxLen, i-start+1)
					start = i
					onMountain, up = true, true
				}
			}
		} else {
			if onMountain {
				if up {
					up = false
				}
			}
		}
	}

	if onMountain && !up {
		maxLen = max(maxLen, len(A)-start)
	}

	if maxLen < 3 {
		maxLen = 0
	}

	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestIt(t *testing.T) {
	testCase:=[]int{8,3,7,3,4,10,1,1,2,8}

	if longestMountain(testCase)!=4 {
		log.Fatal("error result")
	}
}