package hj63dna_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func main() {
	dna, subLen := readInputs()
	fmt.Println(highestRate(dna, subLen))
}

func readInputs() (string, int) {
	var scanner = bufio.NewScanner(os.Stdin)
	lines := []string{}

	for i := 0; i < 2; i++ {
		scanner.Scan()
		lines = append(lines, scanner.Text())
	}

	subLen, _ := strconv.Atoi(lines[1])
	return lines[0], subLen
}

func highestRate(dna string, subLen int) string {
	count := 0

	for i := 0; i < subLen; i++ {
		if dna[i] == 'C' || dna[i] == 'G' {
			count++
		}
	}

	maxCountPos := 0
	maxCount := count

	for i, j := 1, subLen; j < len(dna); {
		if dna[i-1] == 'C' || dna[i-1] == 'G' {
			count--
		}
		if dna[j] == 'C' || dna[j] == 'G' {
			count++
		}
		if count > maxCount {
			maxCount = count
			maxCountPos = i
		}

		i++
		j++
	}

	return dna[maxCountPos : maxCountPos+subLen]
}

func TestIt(t *testing.T) {
	dna := "AACTGTGCACGACCTGA"
	if highestRate(dna, 5) != "GCACG" {
		t.Fatal("bad answer: ", highestRate(dna, 5))
	}
}
