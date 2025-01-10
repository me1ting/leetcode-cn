package countsubstringsthatcanberearrangedtocontainastringii

// 本题的答案与count-substrings-that-can-be-rearranged-to-contain-a-string-i相同

func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	diff := [26]int{} // 记录当前窗口每个字母的数量差值
	for _, c := range word2 {
		diff[c-'a']--
	}
	cnt := 0 // 记录不满足要求的字母的数量
	for _, n := range diff {
		if n < 0 {
			cnt++
		}
	}

	update := func(idx int, n int) {
		diff[idx] += n
		if n == -1 && diff[idx] == -1 {
			cnt++
		} else if n == 1 && diff[idx] == 0 {
			cnt--
		}
	}

	l, r := 0, 0 //[l,r)，初始时l,r相同，表示不包含word1[0]

	var res int64 = 0
	for l < len(word1) {
		for cnt > 0 && r < len(word1) {
			update(int(word1[r]-'a'), 1)
			r++
		}

		if cnt > 0 {
			break
		}

		for cnt == 0 {
			res += int64(len(word1) - r + 1)
			update(int(word1[l]-'a'), -1)
			l++
		}
	}

	return res
}
