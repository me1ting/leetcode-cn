package longest_substring_without_repeating_characters

import "testing"

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
// 思路很简单，维护字符上次出现的位置，出现重复字符时，判断计数位置是否小于重复的字符串的前一个位置+1，
// 如果是，则更新计数位置
// 更新最长不重复子串的长度
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	records := map[byte]int{}
	max := 1
	start := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if j, ok := records[c]; ok {
			len := i - start
			if len > max {
				max = len
			}
			if j+1>start {
				start = j + 1
			}
		}
		records[c] = i
	}

	len := len(s) - start
	if len > max {
		max = len
	}

	return max
}

func TestFunc(t *testing.T) {
	s := "abba"
	result := lengthOfLongestSubstring(s)
	if result != 2 {
		t.Errorf("expected 2 but recived %d\n", result)
	}
}
