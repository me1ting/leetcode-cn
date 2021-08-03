package shortestunsortedcontinuoussubarray

// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
// 可以使用双指针算法：
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
//
//

/*
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r && nums[l] <= nums[l+1] {
		l++
	}

	if l == r {
		return 0
	}

	for l < r && nums[r-1] <= nums[r] {
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
*/

// 题解：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/solution/zui-duan-wu-xu-lian-xu-zi-shu-zu-by-leet-yhlf/
// 方法1排序后逐一两端比较，时间和空间效率都很低
//
// 方法2看似一次遍历，实则是两次遍历，只是放在了一个循环里面；代码很精炼，但硬塞了二次遍历导致难以理解。
// 其基本思想是：
// - 从左往右遍历，记录最大值，从而找到（最短无序连续子数组）右边界。如果更新了最大值，不影响右边界，如果nums[i]<max，说明i属于最短无序连续子数组，可以暂时更新右边界为i。
// - 从右网左遍历，记录最小值，从而找到左边界。如果更新最小值，不影响左边界；如果nums[i]>min，说明i属于最短无序连续子数组，暂时更新左边界为i。
// 由于缓存的问题，运行时间效率略低于双指针算法，但时间复杂度没有区别。

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	max, right := nums[0], -1
	for i := 0; i < n; i++ {
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}
	}
	if right == -1 {
		return 0
	}

	min, left := nums[n-1], n
	for j := n - 1; j >= 0; j-- {
		if nums[j] > min {
			left = j
		} else {
			min = nums[j]
		}
	}
	return right - left + 1
}
