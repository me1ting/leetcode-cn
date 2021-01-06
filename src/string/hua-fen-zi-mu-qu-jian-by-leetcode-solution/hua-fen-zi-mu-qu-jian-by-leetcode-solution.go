package hua_fen_zi_mu_qu_jian_by_leetcode_solution

// https://leetcode-cn.com/problems/partition-labels/solution/hua-fen-zi-mu-qu-jian-by-leetcode-solution/
// 使用两次遍历，查看代码或者参考题解
func partitionLabels(S string) []int {
	last:=[26]int{}
	for i:=0;i< len(S);i++ {
		last[S[i]-'a'] = i
	}

	var part []int
	start,end := 0,0
	for i:=0;i<len(S);i++ {
		end = max(end,last[S[i] - 'a'])
		if end == i {
			part = append(part, end-start+1)
			start = end+1
		}
	}
	return part
}

func  max(i,j int)int {
	if i>j {
		return i
	}
	return j
}