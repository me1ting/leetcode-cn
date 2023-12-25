package hj32mimajiequ

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Println(Answer(input))
}

func Answer(str string) int {
	longest := 0

	for i := 0; i < len(str); i++ {
		j, k := i-1, i+1
		for j >= 0 && k < len(str) && str[j] == str[k] {
			j--
			k++
		}
		if k-(j+1) > longest {
			longest = k - (j + 1)
		}
		if i+1 < len(str) && str[i] == str[i+1] {
			j, k := i-1, i+2
			for j >= 0 && k < len(str) && str[j] == str[k] {
				j--
				k++
			}
			if k-(j+1) > longest {
				longest = k - (j + 1)
			}
		}
	}

	return longest
}
