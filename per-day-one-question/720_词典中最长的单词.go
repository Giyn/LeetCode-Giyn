/*
-------------------------------------
# @Time    : 2022/3/17 0:12:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 720_词典中最长的单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(words))
}

func longestWord(words []string) (ans string) {
	t := &Trie720{}
	for _, word := range words {
		t.Insert(word)
	}
	for _, word := range words {
		if t.Search(word) && (len(word) > len(ans) || len(word) == len(ans) && word < ans) {
			ans = word
		}
	}
	return
}

type Trie720 struct {
	children [26]*Trie720
	isEnd    bool
}

func (t *Trie720) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie720{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie720) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil || !node.children[ch].isEnd {
			return false
		}
		node = node.children[ch]
	}
	return true
}
