package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := readInputs()
	fmt.Println(answer(inputs))
}

func answer(heights []int) int {
	len := len(heights)
	dp1 := make([]int, len) //di1[i]表示以i为终点的最长升序子序列，不包括i
	dp2 := make([]int, len) //di2[i]表示以i为起点的最长降序子序列，不包括i
	for i := 0; i < len; i++ {
		for j := i - 1; j >= 0; j-- {
			if heights[j] < heights[i] {
				temp := 1 + dp1[j]
				if temp > dp1[i] {
					dp1[i] = temp
				}
			}
		}
	}
	for i := len - 1; i >= 0; i-- {
		for j := i + 1; j < len; j++ {
			if heights[j] < heights[i] {
				temp := 1 + dp2[j]
				if temp > dp2[i] {
					dp2[i] = temp
				}
			}
		}
	}

	max := dp1[0] + dp2[0]
	for i := 0; i < len; i++ {
		subLen := dp1[i] + dp2[i]
		if subLen > max {
			max = subLen
		}
	}

	return len - max - 1
}

func readInputs() []int {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for i := 0; i < 2; i++ {
		if scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}

	if len(lines) < 2 {
		panic("error inputs")
	}
	heightStrings := strings.Split(lines[1], " ")

	inputs := []int{}
	for _, str := range heightStrings {
		h, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, int(h))
	}
	return inputs
}
