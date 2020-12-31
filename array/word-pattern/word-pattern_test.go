package word_pattern

import "testing"

// https://leetcode-cn.com/problems/word-pattern/
// 没啥难度
func wordPattern(pattern string, s string) bool {
	p2w :=map[uint8]string{}
	w2p :=map[string]uint8{}
	wordL,wordR:=0,0
	for i:=0;i< len(pattern);i++ {
		p:=pattern[i]

		var word string
		for wordR = wordL;wordR< len(s);wordR++ {
			if s[wordR] == ' ' {
				break
			}
		}
		if wordL<wordR {
			word = s[wordL:wordR]
			wordL = wordR+1
		}else{
			return false
		}

		if w,ok:= p2w[p];ok{
			if w != word {
				return false
			}
		}else{
			if _,ok:=w2p[word];ok {
				return false
			}
			p2w[p] = word
			w2p[word] = p
		}
	}
	if wordR != len(s) {
		return false
	}
	return true
}

func TestWordPattern(t *testing.T) {
	pattern := "abba"
	words:="dog dog dog dog"
	if wordPattern(pattern,words)!=false {
		t.Error("wrong result")
	}
}