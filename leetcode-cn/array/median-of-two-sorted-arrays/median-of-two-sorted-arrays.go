package main

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 这道题其实并不难，只是思路要转过弯：
// 本质和排序的单数组找中位数没什么区别。
// 首先是定义中位数概念：
// len = len(nums1)+len(nums2)，如果len为奇数，值为len/2位置为中位数；否则len/2位置和len/2+1位置的平均为中位数
//
// nums1与nums2交错，但我们可以在nums1中找到一个位置i，计算得到nums2中的位置j(i+j为len的一半)，满足：
//   nums1[i]<=nums[j+1],nums[j]<=num[i+1]
// 在实现过程中，边界条件是比较繁琐的事情。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	half := (len1 + len2 + 1) / 2 //关键，这样如果len是奇数，保证满足条件时中位数位于half中
	isOdd := (len1+len2)%2 == 1

	if len1 > len2 {
		nums1, nums2 = nums2, nums1
		len1, len2 = len2, len1
	}

	l, r := 0, len1-1
	for l <= r {
		i := (l + r) / 2
		j := half - (i + 1) - 1

		if j+1 < len2 && nums1[i] > nums2[j+1] {
			r = i - 1
		} else if i+1 < len1 && nums1[i+1] < nums2[j] {
			l = i + 1
		} else {
			var maxL int
			if j == -1 { // not exist
				maxL = nums1[i]
			} else {
				maxL = max(nums1[i], nums2[j])
			}
			if isOdd {
				return float64(maxL)
			} else {
				var minR int
				if i+1 == len1 { //not exist
					minR = nums2[j+1]
				} else if j+1 == len2 { //not exist
					minR = nums1[i+1]
				} else {
					minR = min(nums1[i+1], nums2[j+1])
				}
				return float64(maxL+minR) / 2
			}
		}
	}

	var j int
	if l == len1 {
		j = half - len1 - 1
	} else {
		j = half - 1
	}
	maxL := nums2[j]
	if isOdd {
		return float64(maxL)
	} else {
		var minR int
		if len1 == 0 {
			minR = nums2[j+1]
		} else if j+1 == len2 {
			minR = nums1[0]
		} else {
			minR = min(nums1[0], nums2[j+1])
		}

		return float64(maxL+minR) / 2
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}

func main() {
	findMedianSortedArrays([]int{3}, []int{1, 2, 4})
}
