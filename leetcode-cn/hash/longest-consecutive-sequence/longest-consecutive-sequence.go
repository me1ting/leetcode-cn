package answer

// https://leetcode-cn.com/problems/longest-consecutive-sequence/
// 使用hashmap来实现数的O(1)查询
// 遍历数组，查找以其起点开始的序列，如果n-1存在于hashmap中，表明其所属序列已经遍历过，因此跳过
// 因此每个数字仅查找一次，算法复杂度为为O(n)
func longestConsecutive(nums []int) int {
	exist := map[int]bool{}
	for _, n := range nums {
		exist[n] = true
	}

	maxLen := 0
	for _, n := range nums {
		if !exist[n-1] {
			i := n + 1
			for exist[i] {
				i++
			}
			curr := i - n
			if curr > maxLen {
				maxLen = curr
			}
		}
	}
	return maxLen
}
