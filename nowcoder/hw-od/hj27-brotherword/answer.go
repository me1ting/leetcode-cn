package hj27brotherword

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	str, list, k := readInputs()
	m, kPointer := answer(str, list, k)
	fmt.Println(m)
	if kPointer != nil {
		fmt.Println(*kPointer)
	}
}

func readInputs() (string, []string, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	slices := strings.Split(input, " ")

	n, _ := strconv.Atoi(slices[0])
	list := slices[1 : n+1]
	str := slices[n+1]
	k, _ := strconv.Atoi(slices[n+2])

	return str, list, k
}

func answer(s string, list []string, k int) (int, *string) {
	checker := NewBrotherChecker(s)
	brothers := []string{}
	for _, item := range list {
		if checker.IsBrother(item) {
			brothers = append(brothers, item)
		}
	}
	if k > len(brothers) {
		return len(brothers), nil
	}
	sort.Strings(brothers)
	return len(brothers), &brothers[k-1]
}

type BrotherChecker struct {
	str    string
	bucket map[byte]int
}

func NewBrotherChecker(s string) *BrotherChecker {
	return &BrotherChecker{
		str:    s,
		bucket: bucket(s),
	}
}

func bucket(s string) map[byte]int {
	b := map[byte]int{}
	for i := 0; i < len(s); i++ {
		b[byte(s[i])]++
	}
	return b
}

func (b *BrotherChecker) IsBrother(s string) bool {
	if s == b.str {
		return false
	}
	if len(s) != len(b.str) {
		return false
	}
	bucket := bucket(s)
	for c, n := range bucket {
		if b.bucket[c] != n {
			return false
		}
	}
	return true
}
