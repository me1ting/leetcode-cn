package answer

import "testing"

func maxRepOpt1(text string) int {
	n := len(text)
	count := map[byte]int{}
	for i := 0; i < n; i++ {
		c := text[i]
		count[c] += 1
	}

	i, j := 0, 1
	res := 1
	for i < n {
		// 求连续相同字符子串
		c := text[i]
		for j < n && text[j] == c {
			j++
		}

		subLen := j - i
		// 更新最值
		// 如果左边或右边可以交换，且有交换对象，那么当前子字符串的长度应当加1
		if (i > 0 || j < n) && count[c] > subLen {
			res = max(res, subLen+1)
		} else {
			res = max(res, subLen)
		}

		// 一种情况是隔一个位置后，存在具有相同字符的另一个子串
		// 最终结果会把这两者连接起来
		k := j + 1
		for k < n && text[k] == c {
			k++
		}

		nextLen := k - j - 1
		if nextLen > 0 {
			totalLen := subLen + nextLen
			if count[c] > totalLen {
				totalLen++
			}
			res = max(res, totalLen)
			if nextLen > 1 {
				i = j + 1
				j = k
				continue
			}
		}
		i = j
		j = i + 1
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestMaxRepOpt1(t *testing.T) {
	texts := []string{"ababa", "aaabaaa", "aaabbaaa", "aaaaa", "abcdef", "bbababaaaa"}
	results := []int{3, 6, 4, 5, 1, 6}

	for i, text := range texts {
		result := results[i]
		ans := maxRepOpt1(text)

		if ans != result {
			t.Errorf("expected %v, but got %v", result, ans)
		}
	}
}
