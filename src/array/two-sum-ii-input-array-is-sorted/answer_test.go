package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	for i, j := 0, n-1; i < j; {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
