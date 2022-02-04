/*
-------------------------------------
# @Time    : 2022/2/2 16:53:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2000_反转单词前缀.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "abcdefd"
	ch := byte('d')
	fmt.Println(reversePrefix(word, ch))
}

func reversePrefix(word string, ch byte) string {
	right := strings.IndexByte(word, ch)
	if right < 0 {
		return word
	}
	s := []byte(word)
	for left := 0; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return string(s)

	//right := -1
	//b := []byte(word)
	//for i := 0; i < len(b); i++ {
	//	if b[i] == ch {
	//		right = i
	//		break
	//	}
	//}
	//if right == -1 {
	//	return word
	//}
	//for i := 0; i < right; i++ {
	//	b[i], b[right] = b[right], b[i]
	//	right--
	//}
	//return string(b)
}
