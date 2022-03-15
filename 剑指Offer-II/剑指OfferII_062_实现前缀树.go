/*
-------------------------------------
# @Time    : 2022/3/15 20:44:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_062_实现前缀树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	t := Constructor062()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("app"))
	t.Insert("app")
	fmt.Println(t.Search("app"))
}

type Trie062 struct {
	children [26]*Trie062
	isEnd    bool
}

func Constructor062() Trie062 {
	return Trie062{}
}

func (this *Trie062) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie062{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie062) SearchPrefix(prefix string) *Trie062 {
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

func (this *Trie062) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie062) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}
