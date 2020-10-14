package find_common_characters

// https://leetcode-cn.com/problems/find-common-characters/
func commonChars(A []string) []string {
	records := make([]int, 26)
	for _, c := range A[0] {
		records[c-'a']++
	}
	for i := 1; i < len(A); i++ {
		newrecords := make([]int, 26)
		for _, c := range A[i] {
			newrecords[c-'a']++
		}
		for j := 0; j < len(records); j++ {
			if records[j] > newrecords[j] {
				records[j] = newrecords[j]
			}
		}
	}
	var result []string
	for i:=0;i<len(records) ;i++  {
		n:=records[i]
		for n > 0 {
			result = append(result, string('a'+i))
			n--
		}
	}
	return result
}
