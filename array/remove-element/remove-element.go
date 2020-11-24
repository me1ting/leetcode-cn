package remove_element

// https://leetcode-cn.com/problems/remove-element/
// 双指针法的典型示例
func removeElement(nums []int, val int) int {
	pos := 0
	for i, n := range nums {
		if n != val {
			if i != pos {
				nums[pos] = n
			}
			pos++
		}
	}
	return pos
}
