/*
-------------------------------------
# @Time    : 2022/3/16 11:44:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_065_最短的单词编码.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	words := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(words))
}

func minimumLengthEncoding(words []string) (ans int) {
	// trie树
	root := &Trie{}
	for _, word := range words {
		root.Insert(root, word)
	}
	return root.findLeaf(root, 0)

	// 哈希
	//mp := map[string]bool{}
	//for _, word := range words {
	//	mp[word] = true
	//}
	//for k := range mp {
	//	// 删除后缀,只保留最长的
	//	for i := 1; i < len(k); i++ {
	//		delete(mp, k[i:])
	//	}
	//}
	//for k := range mp {
	//	ans += len(k) + 1
	//}
	//return
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(root *Trie, word string) {
	for i := len(word) - 1; i >= 0; i-- {
		if root.children[word[i]-'a'] == nil {
			root.children[word[i]-'a'] = &Trie{}
		}
		root = root.children[word[i]-'a']
	}
}

func (t *Trie) findLeaf(root *Trie, length int) (ans int) {
	nodes := root.children
	isLeaf := true
	for i := 0; i < 26; i++ {
		if nodes[i] != nil {
			isLeaf = false
			ans += t.findLeaf(nodes[i], length+1)
		}
	}
	if isLeaf {
		return length + 1
	}
	return
}
