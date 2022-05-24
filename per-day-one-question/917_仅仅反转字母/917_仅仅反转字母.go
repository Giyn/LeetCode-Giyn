/*
-------------------------------------
# @Time    : 2022/2/24 9:49:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 917_仅仅反转字母.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "Test1ng-Leet=code-Q!"
	fmt.Println(reverseOnlyLetters(s))
}

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	left, right := 0, len(b)-1
	for left < right {
		if !unicode.IsLetter(rune(b[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(b[right])) {
			right--
			continue
		}
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
	return string(b)
}
