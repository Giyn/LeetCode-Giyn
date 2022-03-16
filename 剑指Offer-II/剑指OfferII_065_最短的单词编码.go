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
	root := &Trie065{}
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

type Trie065 struct {
	children [26]*Trie065
	isEnd    bool
}

func (this *Trie065) Insert(root *Trie065, word string) {
	for i := len(word) - 1; i >= 0; i-- {
		if root.children[word[i]-'a'] == nil {
			root.children[word[i]-'a'] = &Trie065{}
		}
		root = root.children[word[i]-'a']
	}
}

func (this *Trie065) findLeaf(root *Trie065, length int) (ans int) {
	nodes := root.children
	isLeaf := true
	for i := 0; i < 26; i++ {
		if nodes[i] != nil {
			isLeaf = false
			ans += this.findLeaf(nodes[i], length+1)
		}
	}
	if isLeaf {
		return length + 1
	}
	return
}
