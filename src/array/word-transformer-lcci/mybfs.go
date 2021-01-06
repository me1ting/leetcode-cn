package main

// 使用bfs：
// 基于循环而非递归；bfs时判断adj，该问题可以使用特殊的算法来判断adj。
func findLaddersByMyBfs(beginWord string, endWord string, wordList []string) []string {
	wordMap := map[string]bool{}

	for _, word := range wordList {
		wordMap[word] = true
	}

	if !wordMap[endWord] {
		return nil
	}

	if !wordMap[beginWord] {
		wordMap[beginWord] = true
	}

	queue := map[string]bool{beginWord: true}
	edgeTo := map[string]string{beginWord: beginWord}

	out:
	for len(queue) > 0 {
		nextDepth := map[string]bool{}
		for v,_ := range queue{
			chars := []byte(v)
			for i := 0; i < len(chars); i++ {
				bak := chars[i]
				for c := byte('a'); c <= byte('z'); c++ {
					if bak != c{
						chars[i] = c
						w := string(chars)

						if wordMap[w] {
							if _,ok:=edgeTo[w];!ok{
								nextDepth[w]=true
								edgeTo[w]=v
								if w == endWord {
									break out
								}
							}
						}
					}
				}
				chars[i] = bak
			}
		}
		queue = nextDepth
	}

	if _,ok:=edgeTo[endWord];!ok {
		return nil
	} else {
		paths := make([]string, len(wordMap))
		pos := len(paths)-1
		paths[pos] = endWord
		pos--
		for w := endWord; w != beginWord; {
			v := edgeTo[w]
			paths[pos] = v
			pos--
			w = v
		}
		return paths[pos+1:]
	}
}