package minimumwindowsubstring

import "testing"

// https://leetcode.cn/problems/minimum-window-substring/
// 类似 count-substrings-that-can-be-rearranged-to-contain-a-string-i，使用滑动窗口可以找到一个解
// 通过比较多个解的长度，找到最短的解
func minWindow(s string, t string) string {
	diff := [58]int{} //长度58的数组可以包含'A-Z...a-z'的ASCII内容
	for _, c := range t {
		diff[c-'A']--
	}

	cnt := 0
	for _, n := range diff {
		if n < 0 {
			cnt++
		}
	}

	update := func(idx int, n int) {
		diff[idx] += n
		if n == -1 && diff[idx] == -1 {
			cnt++
		} else if n == 1 && diff[idx] == 0 {
			cnt--
		}
	}

	l, r := 0, 0

	res := ""

	for l < len(s) {
		for cnt > 0 && r < len(s) {
			update(int(s[r]-'A'), 1)
			r++
		}

		//未找到
		if cnt > 0 {
			return res
		}

		for cnt == 0 {
			update(int(s[l]-'A'), -1)
			l++
		}
		match := s[l-1 : r]

		if res == "" {
			res = match
		} else if len(match) < len(res) {
			res = match
		}
	}

	return res
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s      string
		t      string
		expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
	}

	for _, test := range tests {
		result := minWindow(test.s, test.t)
		if result != test.expect {
			t.Errorf("For input s = %s and t = %s, expected %s but got %s", test.s, test.t, test.expect, result)
		}
	}
}
