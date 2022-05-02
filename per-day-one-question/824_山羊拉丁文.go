/*
-------------------------------------
# @Time    : 2022/5/2 14:01:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 824_山羊拉丁文.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "The quick brown fox jumped over the lazy dog"
	fmt.Println(toGoatLatin(sentence))
}

func toGoatLatin(sentence string) string {
	isVowel := func(b byte) bool {
		if b != 'a' && b != 'A' && b != 'e' && b != 'E' && b != 'i' && b != 'I' && b != 'o' && b != 'O' && b != 'u' && b != 'U' {
			return false
		}
		return true
	}
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if !isVowel(words[i][0]) {
			tmp := words[i][0]
			words[i] = words[i][1:]
			words[i] += string(tmp)
		}
		words[i] += "ma"
		words[i] += strings.Repeat("a", i+1)
	}
	return strings.Join(words, " ")
}
