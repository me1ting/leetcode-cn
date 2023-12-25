package main

// https://leetcode-cn.com/problems/k-diff-pairs-in-an-array
// 对于查找指定数，使用map是最佳选择
// 在原始实现中，我实现了一个一次扫描算法，但很复杂，实际运行效率也一半。
// 避免重复遍历的有效手段是遍历map

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	counts := map[int]int{}
	for _,n:=range nums {
		counts[n]++
	}

	count := 0
	for n,c :=range counts{
		if k == 0 {
			if c > 1{
				count++
			}
		}else{
			if counts[n+k] > 0 {
				count++
			}
		}
	}

	return count
}
