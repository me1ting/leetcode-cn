package answer

func removeDuplicates(nums []int) int {
	n := len(nums)
	i, j := 1, 1
	for j < n {
		if nums[j] == nums[j-1] {
			for j < n-1 && nums[j] == nums[j+1] {
				j++
			}
		}
		nums[i] = nums[j]
		i++
		j++
	}
	return i
}
