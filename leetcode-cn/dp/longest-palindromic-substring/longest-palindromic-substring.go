package main

import "fmt"

// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 尝试下动态规划的标准解法

func longestPalindrome(s string) string {
	l := len(s)
	if l <=1 {
		return s
	}

	//dp[j][i]表示长度j+1，以位置i为中心的回文是否存在，如果长度为偶数，i为中心左边的元素
	dp := make([][]bool,l)

	dp[0] = make([]bool,l)
	for i:=0;i<l;i++{
		dp[0][i] = true
	}

	next,nextNext := true,false
	longest:=1

	dp[1] = make([]bool,l)
	for i:=0;i<l-1;i++{
		if s[i] == s[i+1] {
			dp[1][i] = true
			nextNext = true
			longest = 2
		}
	}

	//测试j+1长度的回文是否存在
	j:=2
	for ;(next||nextNext) && j<l;j++{
		curr := next
		next,nextNext = nextNext,false
		dp[j] = make([]bool,l)
		if curr {
			subJ, left, right := j-2, j/2, l-(j+1)/2
			for i := left; i < right; i++ {
				if dp[subJ][i] && s[i-j/2] == s[i+(j+1)/2] {
					dp[j][i] = true
					nextNext = true
					longest = j + 1
				}
			}
		}
	}

	j = longest - 1
	left,right:=j/2,l-(j+1)/2
	for i:=left;i<right;i++{
		if dp[j][i] {
			return s[i-j/2:i-j/2+j+1]
		}
	}
	return ""
}

func main() {
	fmt.Println(longestPalindrome("abcba"))
}