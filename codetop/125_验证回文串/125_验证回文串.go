/*
-------------------------------------
# @Time    : 2022/5/22 13:50:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 125_验证回文串.go
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
	s := "race a car"
	fmt.Println(isPalindromeStr(s))
}

func isPalindromeStr(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l <= r {
		for l < r && !(unicode.IsLetter(rune(s[l])) || unicode.IsDigit(rune(s[l]))) {
			l++
		}
		for l < r && !(unicode.IsLetter(rune(s[r])) || unicode.IsDigit(rune(s[r]))) {
			r--
		}
		if l < r && s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
