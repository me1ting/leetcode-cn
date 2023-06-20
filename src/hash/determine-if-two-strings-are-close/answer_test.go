package determineiftwostringsareclose

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	count1 := make([]int, 26)
	count2 := make([]int, 26)
	for _, c := range word1 {
		count1[c-'a']++
	}

	for _, c := range word2 {
		count2[c-'a']++
	}

	for i := 0; i < 26; i++ {
		if count1[i]*count2[i] == 0 && count1[i]+count2[i] > 0 {
			return false
		}
	}

	sort.Ints(count1)
	sort.Ints(count2)

	for i := 0; i < 26; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}

	return true
}
