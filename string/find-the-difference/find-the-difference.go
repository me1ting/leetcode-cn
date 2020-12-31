package find_the_difference

// https://leetcode-cn.com/problems/find-the-difference/
// 计数法
func findTheDifference(s string, t string) byte {
	counts := [27]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		counts[s[i]-'a']--
		counts[t[i]-'a']++
	}
	counts[t[n]-'a']++
	for i, n := range counts {
		if n > 0 {
			return byte(i) + 'a'
		}
	}
	return 0
}
