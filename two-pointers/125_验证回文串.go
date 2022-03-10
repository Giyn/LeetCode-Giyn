/*
-------------------------------------
# @Time    : 2022/3/10 9:10:33
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
	s := "0P"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !(unicode.IsLetter(rune(s[left])) || unicode.IsDigit(rune(s[left]))) {
			left++
		}
		for left < right && !(unicode.IsLetter(rune(s[right])) || unicode.IsDigit(rune(s[right]))) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
