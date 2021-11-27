/*
-------------------------------------
# @Time    : 2021/10/19 9:01:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 211_添加与搜索单词 - 数据结构设计.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Println(obj.Search("pad"))
	fmt.Println(obj.Search("bad"))
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("b.."))
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	tmp := this
	for _, v := range word {
		v -= 'a'
		if tmp.children[v] == nil {
			tmp.children[v] = &WordDictionary{}
		}
		tmp = tmp.children[v]
	}
	tmp.isEnd = true
}

func dfs211(node *WordDictionary, word string, index int) bool {
	if node == nil {
		return false
	}
	if len(word) == index {
		if node.isEnd {
			return true
		}
		return false
	}
	ch := word[index]
	if ch != '.' {
		ch -= 'a'
		node = node.children[ch]
		return dfs211(node, word, index+1)
	} else {
		res := false
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				res = res || dfs211(node.children[i], word, index+1)
			}
		}
		return res
	}
}

func (this *WordDictionary) Search(word string) bool {
	return dfs211(this, word, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
