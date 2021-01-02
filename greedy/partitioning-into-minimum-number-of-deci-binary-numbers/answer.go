package answer

// https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// 评论说看似贪心，其实脑筋急转弯，没什么问题
func minPartitions(n string) int {
	max := 1
	for i := 0; i < len(n); i++ {
		bit := int(n[i] - '0')
		if bit > max {
			max = bit
		}
	}
	return max
}
