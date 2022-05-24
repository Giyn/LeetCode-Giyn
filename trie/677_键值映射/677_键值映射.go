/*
-------------------------------------
# @Time    : 2022/3/16 15:44:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 677_键值映射.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
}

type Trie struct {
	children [26]*Trie
	val      int
}

type MapSum struct {
	root *Trie
	cnt  map[string]int
}

func Constructor() MapSum {
	return MapSum{
		&Trie{},
		map[string]int{},
	}
}

func (ms *MapSum) Insert(key string, val int) {
	delta := val
	if ms.cnt[key] > 0 {
		delta -= ms.cnt[key]
	}
	ms.cnt[key] = val
	node := ms.root
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
		node.val += delta
	}
}

func (ms *MapSum) Sum(prefix string) int {
	node := ms.root
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	return node.val
}
