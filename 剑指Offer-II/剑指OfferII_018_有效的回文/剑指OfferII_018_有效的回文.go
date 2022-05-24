/*
-------------------------------------
# @Time    : 2022/3/10 8:53:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_018_有效的回文.go
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
	s := "A man, a plan, a canal: Panama"
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
