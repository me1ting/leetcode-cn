package hj20passwodchecker

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const (
	number    = 0x0001
	lower     = 0x0010
	capital   = 0x0100
	other     = 0x1000
	forbidden = -0x0001
)

func main() {

	inputs := readInputs()
	for _, input := range inputs {
		if Check(input) {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}

func readInputs() []string {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}

func Check(password string) bool {
	if len(password) <= 8 {
		return false
	}

	return CheckTypes(password) && CheckCommonSubSequence(password)
}

func CheckTypes(password string) bool {
	types := map[int32]bool{}
	for i := 0; i < len(password); i++ {
		r := password[i]
		types[CharType(r)] = true
	}
	if types[forbidden] || len(types) < 3 {
		return false
	}
	return true
}

func CharType(r byte) int32 {
	if '0' <= r && r <= '9' {
		return number
	}

	if 'a' <= r && r <= 'z' {
		return lower
	}

	if 'A' <= r && r <= 'Z' {
		return capital
	}

	if r != ' ' && r != '\t' && r != '\r' && r != '\n' {
		return other
	}
	return forbidden
}

func CheckCommonSubSequence(password string) bool {
	data := []byte(password)
	stop := len(data) - 5
	for i := 0; i < stop; i++ {
		if bytes.Contains(data[i+3:], data[i:i+3]) {
			return false
		}
	}
	return true
}
