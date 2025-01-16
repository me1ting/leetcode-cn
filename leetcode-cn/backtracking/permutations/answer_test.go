package permutations

// https://leetcode.cn/problems/permutations/
// 实现全排列的思路很简单：以[1,2,3]为例，其结果等于[1]与[2,3]全排列，加上[2]与[1,3]的全排列，加上[3]与[1,2]的全排列。
// 即对于nums，我们依次选出一个元素，并递归对剩余元素做全排列，直到剩余元素长度为1，此时我们得到一个解。
// 对剩余元素做全排列结束后，我们需要放回被选择的元素，以免弄乱的顺序。
func permute(nums []int) [][]int {
	res := [][]int{}
	var backtrack func(remain []int)
	backtrack = func(remain []int) {
		if len(remain) == 1 {
			dst := make([]int, len(nums))
			copy(dst, nums)
			res = append(res, dst)
		}
		for i := 0; i < len(remain); i++ {
			remain[0], remain[i] = remain[i], remain[0]
			backtrack(remain[1:])
			remain[0], remain[i] = remain[i], remain[0]
		}
	}
	backtrack(nums)
	return res
}
