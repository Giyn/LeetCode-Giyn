/*
-------------------------------------
# @Time    : 2022/3/4 9:50:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_005_单词长度的最大乘积.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	fmt.Println(maxProduct(words))
}

func maxProduct(words []string) (ans int) {
	n := len(words)
	wordCode := make([]int, n)
	for i := 0; i < n; i++ {
		for _, ch := range words[i] {
			wordCode[i] |= 1 << (ch - 'a')
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (wordCode[i] & wordCode[j]) == 0 {
				ans = Max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return
}
