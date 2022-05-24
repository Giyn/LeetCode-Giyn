/*
-------------------------------------
# @Time    : 2021/12/29 22:41:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 472_连接词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Println(findAllConcatenatedWordsInADict(words))
}

type trie struct {
	children [26]*trie
	isEnd    bool
}

func (t *trie) insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *trie) dfs(word string) bool {
	if word == "" {
		return true
	}
	node := t
	for i, ch := range word {
		node = node.children[ch-'a']
		if node == nil {
			return false
		}
		if node.isEnd && t.dfs(word[i+1:]) {
			return true
		}
	}
	return false
}

func findAllConcatenatedWordsInADict(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })

	t := &trie{}
	for _, word := range words {
		if word == "" {
			continue
		}
		if t.dfs(word) {
			ans = append(ans, word)
		} else {
			t.insert(word)
		}
	}
	return
}
