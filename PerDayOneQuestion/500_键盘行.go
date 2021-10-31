/*
-------------------------------------
# @Time    : 2021/10/31 2:22:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 500_键盘行.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

func findWords(words []string) (ans []string) {
	const rowIdx = "12210111011122000010020202"
next:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, ch := range word[1:] {
			if rowIdx[unicode.ToLower(ch)-'a'] != idx {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return
}
