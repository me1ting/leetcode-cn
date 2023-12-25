package lettercombinationsofaphonenumber

var digit2chars = []string{
	"",
	"",
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	size := 1
	for i := 0; i < len(digits); i++ {
		size *= 4
	}
	combinations = make([]string, 0, size)
	backtrack(digits, 0, make([]byte, 0, len(digits)))
	return combinations
}

func backtrack(digits string, index int, combination []byte) {
	if index == len(digits) {
		combinations = append(combinations, string(combination))
		return
	}
	digit := digits[index] - '0'
	chars := digit2chars[digit]

	for _, char := range chars {
		backtrack(digits, index+1, append(combination, byte(char)))
	}
}
