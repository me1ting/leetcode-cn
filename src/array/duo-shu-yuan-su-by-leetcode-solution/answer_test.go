package answer

func majorityElement(nums []int) int {
	curr := nums[0]
	vote := 1
	for i := 1; i < len(nums); i++ {
		if vote == 0 {
			curr = nums[i]
			vote++
			continue
		}

		if nums[i] == curr {
			vote++
		} else {
			vote--
		}
	}
	return curr
}
