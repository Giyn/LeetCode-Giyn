/*
-------------------------------------
# @Time    : 2022/3/16 0:59:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 648_单词替换.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
}

func replaceWords(dictionary []string, sentence string) string {
	// 前缀树
	trie := Constructor()
	for _, root := range dictionary {
		trie.Insert(root)
	}
	sentences := strings.Split(sentence, " ")
	for i, s := range sentences {
		sentences[i] = trie.MinPrefix(s)
	}
	return strings.Join(sentences, " ")

	// 调库
	//sort.Strings(dictionary)
	//sentences := strings.Split(sentence, " ")
	//for i, s := range sentences {
	//	for _, root := range dictionary {
	//		if strings.HasPrefix(s, root) {
	//			sentences[i] = root
	//			break
	//		}
	//	}
	//}
	//return strings.Join(sentences, " ")
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) MinPrefix(s string) string {
	var res []byte
	node := t
	for i, ch := range s {
		ch -= 'a'
		if node.children[ch] == nil {
			return s
		}
		res = append(res, s[i])
		node = node.children[ch]
		if node.isEnd {
			return string(res)
		}
	}
	return s
}
