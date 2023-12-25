package answer

import "testing"

type Trie struct {
	next  []*Trie
	exist bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next: make([]*Trie, 27),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		index := c - 'a'
		if node.next[index] == nil {
			trie := Constructor()
			node.next[index] = &trie
		}
		node = node.next[index]
	}
	node.exist = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	i := 0
	for ; i < len(word); i++ {
		c := word[i]
		index := c - 'a'
		if node.next[index] == nil {
			break
		}
		node = node.next[index]
	}
	if i == len(word) && node.exist {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	i := 0
	for ; i < len(prefix); i++ {
		c := prefix[i]
		index := c - 'a'
		if node.next[index] == nil {
			break
		}
		node = node.next[index]
	}
	if i == len(prefix) {
		if node.exist {
			return true
		}

		for _, nextNode := range node.next {
			if nextNode != nil {
				return true
			}
		}
	}
	return false
}

func TestIt(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	val := trie.Search("apple") // 返回 true
	if val != true {
		t.Errorf("error")
		return
	}
	val = trie.Search("app") // 返回 false
	if val != false {
		t.Errorf("error")
		return
	}
	val = trie.StartsWith("app") // 返回 true
	if val != true {
		t.Errorf("error")
		return
	}
	trie.Insert("app")
	val = trie.Search("app") // 返回 true
	if val != true {
		t.Errorf("error")
		return
	}
}
