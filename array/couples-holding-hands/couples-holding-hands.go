package couples_holding_hands

// https://leetcode-cn.com/problems/couples-holding-hands/
// 因为只有两两交换而没有其它限制
// 任意交换满足一对情侣在n~n+1（n为偶数）位置上，就是最佳选择（直觉但没法证明），因此使用贪婪算法即可
func minSwapsCouples(row []int) int {
	count := 0
	for i := 0; i < len(row); i += 2 {
		they:= row[i]
		var right int
		if they%2 == 0 {
			right = they + 1
		}else{
			right = they - 1
		}
		if row[i+1] != right {
			for j := 0; j < len(row); j++ {
				if row[j] == right {
					row[i+1], row[j] = row[j], row[i+1]
					count++
				}
			}
		}
	}
	return count
}