package hj26stringsort

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	input := readInputs()
	bs := []byte(input)
	CustomSort(bs)
	fmt.Println(string(bs))
}

func readInputs() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func CustomSort(s []byte) {
	unalphabetIndexs := map[int]bool{}
	alphabets := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		r := s[i]
		if 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' {
			alphabets = append(alphabets, r)
		} else {
			unalphabetIndexs[i] = true
		}
	}
	sort.Stable(CustomSortBytes(alphabets))

	for i, j := 0, 0; j < len(alphabets); j++ {
		for unalphabetIndexs[i] {
			i++
		}

		s[i] = alphabets[j]
		i++
	}
}

type CustomSortBytes []byte

// Less implements sort.Interface.
func (c CustomSortBytes) Less(i int, j int) bool {
	loweri := unicode.ToLower(rune(c[i]))
	lowerj := unicode.ToLower(rune(c[j]))
	return loweri < lowerj
}

// Swap implements sort.Interface.
func (c CustomSortBytes) Swap(i int, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CustomSortBytes) Len() int {
	return len(c)
}

var _ sort.Interface = CustomSortBytes([]byte{})
