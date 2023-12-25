package stringcompression

import "fmt"

func compress(chars []byte) int {
	pos := 0
	curr := chars[0]
	n := 1

	add := func() {
		chars[pos] = curr
		pos++
		if n > 1 {
			nText := fmt.Sprintf("%d", n)
			_ = append(chars[:pos], nText...)
			pos += len(nText)
		}
	}

	for i := 1; i < len(chars); i++ {
		char := chars[i]
		if char == curr {
			n++
		} else {
			add()
			curr = char
			n = 1
		}
	}
	add()
	return pos
}
