package counttherepetitions_test

import "testing"

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	if n1%n2 == 0 {
		n1 = n1 / n2
		n2 = 1
	}

	count := 0
	lenStr2 := n2 * len(s2)
	k, realK := 0, 0
	for i := 0; i < n1; i++ {
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[k] {
				k++
				realK++
				if k == len(s2) {
					k = 0
					if realK == lenStr2 {
						realK = 0
						count++
					}
				}
			}
		}
	}

	return count
}

func TestIt(t *testing.T) {
	s1, s2 := "abc", "ab"
	n1, n2 := 4, 2
	if getMaxRepetitions(s1, n1, s2, n2) != 2 {
		t.Fatal("bad answer: ", getMaxRepetitions(s1, n1, s2, n2))
	}
}
