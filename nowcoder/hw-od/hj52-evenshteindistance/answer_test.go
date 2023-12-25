package hj52evenshteindistance_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func main() {
	inputs := readInputs()
	fmt.Println(distance(inputs[0], inputs[1]))
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

func distance(a, b string) int {
	// dp [i][j] 表示a[0:i],b[0:j]的编辑距离
	dp := make([][]int, len(a)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		dp[i][0] = i
	}

	for i := 0; i <= len(b); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			dp[i][j] = min(dp[i][j], min(dp[i-1][j]+1, dp[i][j-1]+1))
		}
	}

	return dp[len(a)][len(b)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestIt(t *testing.T) {
	a := "zikwvkijajpkaicihcffiemzexmwjjlyrylxcuoewdmpivudhmgkuodjaurazdjnlgtpwz"
	b := "wpnmubqfsnmapqpufmmsphqehjplwjkqspnnpywsvvjilxbcfsrygbelquaalenvkruyltiwqcpdrxgstywaja"

	if distance(a, b) != 73 {
		t.Fatal("bad answer", distance(a, b))
	}
}
