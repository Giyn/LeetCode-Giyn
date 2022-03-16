/*
-------------------------------------
# @Time    : 2022/3/16 9:44:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_064_神奇的字典.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	md := Constructor064()
	md.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(md.Search("hello"))
	fmt.Println(md.Search("hhllo"))
	fmt.Println(md.Search("hell"))
	fmt.Println(md.Search("leetcoded"))
}

type MagicDictionary struct {
	children [26]*MagicDictionary
	isWord   bool
}

func Constructor064() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &MagicDictionary{}
		}
		node = node.children[ch]
	}
	node.isWord = true
}

func (this *MagicDictionary) IsExist(word string) bool {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isWord
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.Insert(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	node := this
	b := []byte(searchWord)
	for i := 0; i < len(b); i++ {
		oriB := b[i]
		// 替换
		for j := 0; j < 26; j++ {
			if byte('a'+j) == oriB {
				continue
			}
			b[i] = byte('a' + j)
			if node.IsExist(string(b[i:])) {
				return true
			}
		}
		b[i] = oriB // 恢复
		// 没有匹配到能替换一次的情况
		if node.children[b[i]-'a'] == nil {
			return false
		}
		node = node.children[b[i]-'a'] // 结点移动,继续搜索
	}
	return false
}
