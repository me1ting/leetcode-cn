package hj71tokens_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func main() {
	a, b := readInputs()
	a, b = strings.ToLower(a), strings.ToLower(b)
	fmt.Println(test([]byte(a), []byte(b)))
}

func readInputs() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for i := 0; i < 2; i++ {
		scanner.Scan()
		inputs = append(inputs, scanner.Text())
	}
	return inputs[0], inputs[1]
}

func test(exp []byte, str []byte) bool {
	if len(exp) == 0 {
		if len(str) == 0 {
			return true
		}
		return false
	}

	i, j := 0, 0
	for i < len(exp) && j < len(str) {
		if exp[i] == '*' {
			for i < len(exp) && exp[i] == '*' {
				i++
			}
			for m := j; m <= len(str); m++ {
				if test(exp[i:], str[m:]) {
					return true
				}
				if m < len(str) && !isASCIIorDigit(str[m]) {
					break
				}
			}
			return false
		}

		if !(exp[i] == '?' && isASCIIorDigit(str[j]) || exp[i] == str[j]) {
			return false
		}
		i++
		j++
	}
	if i != len(exp) || j != len(str) {
		return false
	}
	return true
}

func isASCIIorDigit(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= '0' && b <= '9'
}

func TestIt(t *testing.T) {
	exp, str := "t?t*1*.*", "txt12.xls"

	if !test([]byte(exp), []byte(str)) {
		t.Fatal("bad answer")
	}
}
