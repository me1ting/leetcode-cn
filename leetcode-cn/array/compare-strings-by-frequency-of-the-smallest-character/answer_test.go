package comparestringsbyfrequencyofthesmallestcharacter

import "testing"

func f(s string) int {
	c := 'z'
	count := 0

	for _, char := range s {
		if char < c {
			c = char
			count = 1
		} else if char == c {
			count++
		}
	}
	return count
}

func numSmallerByFrequency(queries []string, words []string) []int {
	count := make([]int, 12) // 预计的最大长度为10，这里取12是为了优化算法，不需要考虑最末尾的元素
	for _, s := range words {
		count[f(s)]++
	}

	for i := 9; i >= 0; i-- { // 如果长度刚好为10，需要考虑末尾元素
		count[i] = count[i] + count[i+1]
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = count[f(query)+1]
	}
	return ans
}

func TestF(t *testing.T) {
	text := "zaaaz"
	if f(text) != 3 {
		t.Errorf("bad f")
	}
}
