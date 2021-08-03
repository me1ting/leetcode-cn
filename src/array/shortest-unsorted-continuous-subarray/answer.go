package shortestunsortedcontinuoussubarray

// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
// 可以使用贪婪算法：
// 我们假设数组的左边存在一个有序子数列，数组右边存在一个有序子数列，我们称左有序子数列的右边界（包含）、右有序子数列的左边界（包含）分别为（l,r）：
// - l==len-1，整个数组有序，返回0
// - l==0&&r==len-1，整个输入无序，返回len
// - 一般情况，我们可以称子数列nums[l:r+1]严格无序：左右不存在有序子数列
//
// 下面对于一般情况进行说明：
// 我们需要的结果是一个非严格无序子数列的长度：
// - 左边界new_l应当满足：nums[new_l]>min(nums[l:r+1])
// - 右边界new_r应当满足：nums[new_r]<max(nums[l:r+1])
//
// 也就是我们需要在nums[l,r+1]的基础上找到其最值并往两边扩展，直到满足最短无序连续子数组条件。
//
// 如：1,5,6,7,8,9,2,3,4，nums[l]=9，nums[new_l]=5
// 如：1,2,3,7,4,5,6,8，nums[r]=4，nums[new_r]=6

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r && nums[l] < nums[l+1] {
		l++
	}

	if l == r {
		return 0
	}

	for l < r && nums[r-1] < nums[r] {
		r--
	}

	if l == 0 && r == len(nums)-1 {
		return n
	}

	min, max := nums[l], nums[l]
	for i := l; i <= r; i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	for l >= 0 && nums[l] > min {
		l--
	}
	l++

	for r < n && nums[r] < max {
		r++
	}
	r--

	return r - l + 1
}
