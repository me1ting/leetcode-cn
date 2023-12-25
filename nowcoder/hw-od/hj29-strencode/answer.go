package hj29strencode

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputs := readInputs()
	result := Answer(inputs)
	fmt.Println(result[0])
	fmt.Println(result[1])
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

func Answer(inputs []string) []string {
	return []string{encode(inputs[0]), decode(inputs[1])}
}

var encodeMap = map[byte]byte{}
var decodeMap = map[byte]byte{}
var lowerToCaptial = 'A' - 'a'

func init() {
	for i := 'a'; i < 'z'; i++ {
		encodeMap[byte(i)] = byte(i + 1 + lowerToCaptial)
	}
	for i := 'A'; i < 'Z'; i++ {
		encodeMap[byte(i)] = byte(i + 1 - lowerToCaptial)
	}
	for i := '0'; i < '9'; i++ {
		encodeMap[byte(i)] = byte(i + 1)
	}
	encodeMap[byte('z')] = byte('A')
	encodeMap[byte('Z')] = byte('a')
	encodeMap[byte('9')] = byte('0')

	for i := 'z'; i > 'a'; i-- {
		decodeMap[byte(i)] = byte(i - 1 + lowerToCaptial)
	}
	for i := 'Z'; i > 'A'; i-- {
		decodeMap[byte(i)] = byte(i - 1 - lowerToCaptial)
	}
	for i := '9'; i > '0'; i-- {
		decodeMap[byte(i)] = byte(i - 1)
	}
	decodeMap[byte('a')] = byte('Z')
	decodeMap[byte('A')] = byte('z')
	decodeMap[byte('0')] = byte('9')
}

func encode(str string) string {
	bs := []byte(str)
	for i := 0; i < len(bs); i++ {
		if replace, ok := encodeMap[bs[i]]; ok {
			bs[i] = replace
		}
	}
	return string(bs)
}

func decode(str string) string {
	bs := []byte(str)
	for i := 0; i < len(bs); i++ {
		if replace, ok := decodeMap[bs[i]]; ok {
			bs[i] = replace
		}
	}
	return string(bs)
}
