package contains_duplicate

// https://leetcode-cn.com/problems/contains-duplicate
// hash table,sort 均可以实现目的
func containsDuplicate(nums []int) bool {
	record := map[int]bool{}
	for _, v := range nums {
		if _, ok := record[v]; ok {
			return true
		}
		record[v] = true
	}
	return false
}
