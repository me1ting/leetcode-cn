package searchsuggestionssystem

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := NewNode()
	for _, word := range products {
		trie.Insert(word)
	}

	n := len(searchWord)
	ans := make([][]string, n)
	node := trie
	for i, ch := range searchWord {
		pos := ch - 'a'
		node = node.child[pos]
		if node == nil {
			break
		}
		ans[i] = node.suggest
		sort.Strings(ans[i])
	}
	return ans
}

type Node struct {
	child   []*Node
	suggest []string
}

func NewNode() *Node {
	return &Node{
		child:   make([]*Node, 27),
		suggest: make([]string, 0, 3),
	}
}

func (n *Node) Insert(word string) {
	node := n
	for _, ch := range word {
		pos := ch - 'a'
		if node.child[pos] == nil {
			node.child[pos] = NewNode()
		}
		node = node.child[pos]
		node.UpdateSuggest(word)
	}
}

func (n *Node) UpdateSuggest(word string) {
	if len(n.suggest) < 3 {
		n.suggest = append(n.suggest, word)
		return
	}

	maxPos := 0
	for i := 1; i < 3; i++ {
		if n.suggest[i] > n.suggest[maxPos] {
			maxPos = i
		}
	}
	if n.suggest[maxPos] > word {
		n.suggest[maxPos] = word
	}
}
