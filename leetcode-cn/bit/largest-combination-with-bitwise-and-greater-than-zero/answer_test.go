package largestcombinationwithbitwiseandgreaterthanzero

// https://leetcode.cn/problems/largest-combination-with-bitwise-and-greater-than-zero/
// 一个很简单的道理：对于结果集`res`([]int类型)，存在第 n 位使得结果集中的所有整数都满足该位上的值为1。
func largestCombination(candidates []int) int {
	cnt := [32]int{}
	for _, n := range candidates {
		for i := 0; i < 32; i++ {
			if n&(1<<i) > 0 {
				cnt[i]++
			}
		}
	}

	max := 0
	for _, c := range cnt {
		if c > max {
			max = c
		}
	}

	return max
}

// 优化1：
// 考虑题目规定`1 <= candidates[i] <= 10^7`，因此只需要判断低24位。
//
// 优化2：
// 对比以下官方题解，两者思路相同，但是性能差距却很大：以上答案的运行时间为30ms，而官方题解只需要3ms
// 这是因为以上答案交替读写内存len(candidates)*(24*2+1)次，而官方题解只读内存len(candidates)*24次。
func largestCombination2(candidates []int) int {
	// 计算从低到高第 k 个二进制位数值为 1 的元素个数
	maxlen := func(k int) int {
		res := 0
		for _, num := range candidates {
			if num&(1<<k) != 0 {
				res++
			}
		}
		return res
	}

	res := 0
	for i := 0; i < 24; i++ {
		// 遍历二进制位
		res = max(res, maxlen(i))
	}
	return res
}
