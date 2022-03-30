/*
-------------------------------------
# @Time    : 2022/3/30 23:48:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 557_反转字符串中的单词 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	reverseString := func(s string) string {
		b := []byte(s)
		left, right := 0, len(s)-1
		for left < right {
			b[left], b[right] = b[right], b[left]
			left++
			right--
		}
		return string(b)
	}
	words := strings.Split(s, " ")
	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}
	return strings.Join(words, " ")
}
