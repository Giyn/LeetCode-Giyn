/*
-------------------------------------
# @Time    : 2022/1/27 0:14:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2047_句子中的有效单词数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sentence := "cat and  dog"
	fmt.Println(countValidWords(sentence))
}

func countValidWords(sentence string) (ans int) {
	tokens := strings.Fields(sentence)
outer:
	for _, token := range tokens {
		hasHyphen := false
		for i, ch := range token {
			if unicode.IsDigit(ch) || strings.ContainsRune("!.,", ch) && i < len(token)-1 {
				continue outer
			}
			if ch == '-' {
				if hasHyphen || i == 0 || i == len(token)-1 || !unicode.IsLower(rune(token[i-1])) || !unicode.IsLower(rune(token[i+1])) {
					continue outer
				}
				hasHyphen = true
			}
		}
		ans++
	}
	return
}
