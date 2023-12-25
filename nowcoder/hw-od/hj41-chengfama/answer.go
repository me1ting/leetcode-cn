package hj41chengfama

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := Answer2(readInputs())
	fmt.Println(a)
}

func readInputs() (int, []int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	types, _ := strconv.Atoi(lines[0])
	weightStrings := strings.Split(lines[1], " ")
	weights := []int{}
	for _, s := range weightStrings {
		w, _ := strconv.Atoi(s)
		weights = append(weights, w)
	}

	countStrings := strings.Split(lines[2], " ")
	counts := []int{}
	for _, s := range countStrings {
		c, _ := strconv.Atoi(s)
		counts = append(counts, c)
	}
	return types, weights, counts
}

func Answer1(types int, weights []int, counts []int) int {
	var results = map[int]bool{}
	var scan func(t int, sum int)

	scan = func(t int, sum int) {
		for i := 0; i <= counts[t]; i++ {
			w := sum + i*weights[t]
			if t+1 == types {
				results[w] = true
			} else {
				scan(t+1, w)
			}
		}
	}

	scan(0, 0)

	return len(results)
}

func Answer2(types int, weights []int, counts []int) int {
	records := map[int]bool{0: true}

	for i := 0; i < types; i++ {
		newRecords := map[int]bool{}
		for j := 0; j <= counts[i]; j++ {
			for n := range records {
				newRecords[n+j*weights[i]] = true
			}
		}
		records = newRecords
	}

	return len(records)
}
