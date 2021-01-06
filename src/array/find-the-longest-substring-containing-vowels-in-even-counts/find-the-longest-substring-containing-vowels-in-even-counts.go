package main

// https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts
// 此题的关键在于，要想通：https://leetcode-cn.com/problems/longest-palindromic-substring
// 状态压缩：可以只记录每个元音字母的统计个数是奇数还是偶数，称为状态，这样只需要记录32种状态出现的前后位置
// 相同状态的两个位置i,j，字符串(i,j](除000000的特殊情况)是元音字母为偶数的子字符串
func findTheLongestSubstring(s string) int {
	status:=0b00000
	masks:=map[rune]int{'a':0b00001,'e':0b00010,'i':0b00100,'o':0b01000,'u':0b10000}

	begin,end:=map[int]int{status:-1},map[int]int{}

	for i,c:=range s {
		switch c {
		case 'a','e','i','o','u':
			end[status]=i-1
			status ^= masks[c]
			if _,ok:=begin[status];!ok{
				begin[status] = i
			}
		}
	}
	end[status] = len(s)-1

	longest := 0
	for theStatus,r := range end {
		if length := r-begin[theStatus];length > longest {
			longest = length
		}
	}

	return longest
}