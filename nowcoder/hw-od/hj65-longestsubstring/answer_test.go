package hj65longestsubstring_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func main() {
	inputs := readInputs()
	fmt.Println(longestSubString(inputs[0], inputs[1]))
}

func readInputs() []string {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for i := 0; i < 2; i++ {
		scanner.Scan()
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}

func longestSubString(a, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	//dp[i][j]表示，a[0:i],b[0:j]，以a[i-1],b[j-1]为结尾时的最长公共子串长度
	dp := make([][]int, len(a)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
		}
	}

	longest := 1
	endA := 0
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if dp[i][j] > longest {
				longest = dp[i][j]
				endA = i
			}
		}
	}
	return a[endA-longest : endA]
}

func TestIt(t *testing.T) {
	a, b := "nvlrzqcjltmrejybjeshffenvkeqtbsnlocoyaokdpuxutrsmcmoearsgttgyyucgzgcnurfbubgvbwpyslaeykqhaaveqxijc", "wkigrnngxehuiwxrextitnmjykimyhcbxildpnmrfgcnevjyvwzwuzrwvlomnlogbptornsybimbtnyhlmfecscmojrxekqmj"
	if longestSubString(a, b) != "gcn" {
		t.Fatal("bad answer: ", longestSubString(a, b))
	}
}
