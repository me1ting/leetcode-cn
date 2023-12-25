package main

import "fmt"

// 使用DFS来计算有向图的单点路径
// DFS实现参考《算法４》
// 算法执行效率一半，对比其它实现，发现主要原因是练习太少，导致实现过于死板
func findLadders(beginWord string, endWord string, wordList []string) []string {
	l := len(wordList)
	start, end := l, -1

	for i, word := range wordList {
		if word == endWord {
			end = i
		}
		if word == beginWord {
			start = i
		}
	}

	if end == -1 {
		return []string{}
	}

	pointCount := l
	if start == l {
		pointCount = l + 1
	}

	adj := make([][]int, pointCount) //adj[i]表示v(i,j)的集合
	for i := 0; i < len(adj); i++ {
		adj[i] = []int{}
	}

	if start == l {
		for i, word := range wordList {
			if oneCharDiff(beginWord, word) {
				adj[start] = append(adj[start], i)
			}
		}
	}

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if oneCharDiff(wordList[i], wordList[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	//dfs求单点路径
	edgeTo := make([]int, pointCount)
	marked := make([]bool, pointCount)

	//如何编写递归调用的内部函数
	//https://stackoverflow.com/questions/28099441/define-a-recursive-function-within-a-function-in-go
	var dfs func(v int)
	var findEnd bool // 如果搜索到end，停止dfs
	dfs = func(v int) {
		marked[v] = true
		if findEnd {
			return
		}
		for _, w := range adj[v] {
			if !marked[w] {
				edgeTo[w] = v

				if w == end {
					findEnd = true
				}

				dfs(w)
			}
		}
	}

	dfs(start)

	if !marked[end] {
		return []string{}
	} else {
		paths := make([]string, pointCount)
		pos := pointCount - 1
		paths[pos] = endWord
		pos--
		for w := end; w != start; {
			v := edgeTo[w]
			var vWord string
			if v == l {
				vWord = beginWord
			} else {
				vWord = wordList[v]
			}
			paths[pos] = vWord
			pos--
			w = v
		}
		return paths[pos+1:]
	}
}

func oneCharDiff(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	diff := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}

	return diff == 1
}

func main() {
	//val := findLadders("a", "c", []string{"a", "b", "c"})
	//fmt.Println(val)
	fmt.Println(findLaddersByMyBfs("hit","cog",[]string{"hit","hot","dot","lot","log","cog"}))
}
