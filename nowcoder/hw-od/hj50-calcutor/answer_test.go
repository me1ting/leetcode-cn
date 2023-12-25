package hj50sizeyunsuan_test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"unicode"
)

func main() {
	tokens := readInputs()
	tokens = reverse(tokens)
	fmt.Println(compute(tokens))
}

func readInputs() []any {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return parse(line)
}

func parse(line string) []any {
	stack := []any{}
	onDigit := false
	digitStart := 0
	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(rune(line[i])) {
			if onDigit {
				continue
			}
			onDigit = true
			digitStart = i
			continue
		}

		if onDigit {
			n, _ := strconv.Atoi(line[digitStart:i])
			stack = append(stack, n)
			onDigit = false
		}
		operator := rune(line[i])
		if (operator == '+' || operator == '-') && len(stack) >= 1 {
			if t, ok := stack[len(stack)-1].(rune); ok {
				if isLeftBracket(t) {
					stack = append(stack, 0)
				}
			}
		}
		stack = append(stack, operator)
	}

	if onDigit {
		n, _ := strconv.Atoi(line[digitStart:])
		stack = append(stack, n)
	}

	return stack
}

func reverse(tokens []any) []any {
	output := []any{}     // 使用一个栈记录输出
	operators := []rune{} // 使用一个栈记录操作符

	for _, t := range tokens {
		switch t := t.(type) {
		case int:
			output = append(output, t) // 如果是操作数，放入输出栈中
		case rune:
			if isRightBracket(t) { // 如果是右括号
				i := len(operators) - 1
				for ; i >= 0 && !isLeftBracket(operators[i]); i-- { // 弹出所有操作符，直到遇到匹配的左括号
					output = append(output, operators[i])
				}
				operators = operators[:i]
			} else if isLeftBracket(t) { //如果是左括号，直接入符号栈
				operators = append(operators, t)
			} else {
				i := len(operators) - 1
				tPriority := priority(t)
				for ; i >= 0 && !isLeftBracket(operators[i]) && priority(operators[i]) >= tPriority; i-- { // 弹出所有非左括号且优先级高于或等于当前符号的操作符
					output = append(output, operators[i])
				}
				operators = operators[:i+1]
				operators = append(operators, t) //然后将符号压入符号栈中
			}
		}
		//printTokens(output)
		//printOperators(operators)
	}

	for i := len(operators) - 1; i >= 0; i-- {
		output = append(output, operators[i])
	}

	return output
}

func isLeftBracket(r rune) bool {
	return r == '[' || r == '{' || r == '('
}

func isRightBracket(r rune) bool {
	return r == ']' || r == '}' || r == ')'
}

func priority(r rune) int {
	if isLeftBracket(r) { //|| isRightBracket(r)这里不判断右括号是因为性能而没有必要
		return 3
	}
	if r == '*' || r == '/' {
		return 2
	}
	return 1
}

func compute(tokens []any) int {
	stack := []int{}

	for _, t := range tokens {
		switch t := t.(type) {
		case int:
			stack = append(stack, t)
		case rune:
			top := len(stack) - 1
			stack[top-1] = operate(stack[top-1], stack[top], t)
			stack = stack[:top]
		}
	}

	return stack[0]
}

func operate(i, j int, o rune) int {
	switch o {
	case '+':
		return i + j
	case '-':
		return i - j
	case '*':
		return i * j
	case '/':
		return i / j
	}
	return 0
}

func TestIt(t *testing.T) {
	line := "3+2*{1+2*[-4/(8-6)+7]}"
	tokens := parse(line)
	tokens = reverse(tokens)
	val := compute(tokens)

	if val != 25 {
		t.Fatal("bad answer")
	}
}

func printTokens(tokens []any) {
	for _, t := range tokens {
		switch t := t.(type) {
		case rune:
			fmt.Printf("%c ", t)
		case int:
			fmt.Printf("%d ", t)
		}
	}
	fmt.Println()
}

func printOperators(ops []rune) {
	for _, t := range ops {
		fmt.Printf("%c ", t)
	}
	fmt.Println()
}
