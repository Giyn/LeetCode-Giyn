/*
-------------------------------------
# @Time    : 2022/3/16 12:35:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_066_单词之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	ms := Constructor066()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
}

type Trie066 struct {
	children [26]*Trie066
	val      int
}

type MapSum struct {
	root *Trie066
	cnt  map[string]int
}

func Constructor066() MapSum {
	return MapSum{
		&Trie066{},
		map[string]int{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	delta := val
	if this.cnt[key] > 0 {
		delta -= this.cnt[key]
	}
	this.cnt[key] = val
	node := this.root
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie066{}
		}
		node = node.children[ch]
		node.val += delta
	}
}

func (this *MapSum) Sum(prefix string) int {
	node := this.root
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	return node.val
}
