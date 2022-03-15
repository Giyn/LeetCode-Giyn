/*
-------------------------------------
# @Time    : 2022/3/15 21:46:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 208_实现 Trie (前缀树).go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	t := Constructor208()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("app"))
	t.Insert("app")
	fmt.Println(t.Search("app"))
}

type Trie208 struct {
	children [26]*Trie208
	isEnd    bool
}

func Constructor208() Trie208 {
	return Trie208{}
}

func (this *Trie208) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie208{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie208) SearchPrefix(prefix string) *Trie208 {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie208) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie208) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}
