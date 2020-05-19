package main

// https://leetcode-cn.com/problems/word-transformer-lcci/solution/golang-20ms-shuang-xiang-bfs-by-minato-aqua/
// 能理解，性能更高的原因：
// 1.避免通过字符比较来判断所有字符串是否是其邻接点
// 2.双向bfs
// 3.bfs时判断邻接关系，而非先生成邻接表
func findLaddersBy2BFS(beginWord string, endWord string, wordList []string) []string {
	path := map[string][]string{beginWord: []string{beginWord}, endWord: []string{endWord}}
	startQueue := map[string]bool{beginWord: true}
	endQueue := map[string]bool{endWord: true}
	wordDict := make(map[string]bool)
	for _, word := range wordList {
		wordDict[word] = true
	}
	if !wordDict[endWord] {
		return nil
	}
	// 用来记录是否为正向遍历，因为反向遍历时记录新路径的方向与正向遍历相反
	isBackward := false

	for len(startQueue) > 0 && len(endQueue) > 0 {
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
			isBackward = !isBackward
		}
		tempQueue := make(map[string]bool)
		for word := range startQueue {
			chars := []byte(word)
			for i := 0; i < len(chars); i++ {
				old := chars[i]
				for j := byte('a'); j <= 'z'; j++ {
					chars[i] = j
					newWord := string(chars)
					if endQueue[newWord] {//如果该word存在于另一条路径中，说明找到
						startLen, endLen := len(path[word]), len(path[newWord])
						res := make([]string, startLen+endLen)
						if isBackward {
							copy(res, path[newWord])
							copy(res[endLen:], path[word])
						} else {
							copy(res, path[word])
							copy(res[startLen:], path[newWord])
						}
						return res
					}
					if j != old && wordDict[newWord] && len(path[newWord]) == 0 {
						path[newWord] = make([]string, len(path[word])+1)
						if isBackward {
							path[newWord][0] = newWord
							copy(path[newWord][1:], path[word])
						} else {
							copy(path[newWord], path[word])
							path[newWord][len(path[newWord])-1] = newWord
						}
						tempQueue[newWord] = true
					}
				}
				chars[i] = old
			}
		}
		startQueue = tempQueue
	}

	return nil
}
