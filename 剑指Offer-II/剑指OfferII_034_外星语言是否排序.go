/*
-------------------------------------
# @Time    : 2022/3/12 17:39:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_034_外星语言是否排序.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}

func isAlienSorted(words []string, order string) bool {
	mp := make(map[byte]int, 26)
	for i := 0; i < len(order); i++ {
		mp[order[i]] = i
	}
outer:
	for i := 1; i < len(words); i++ {
		for j := 0; j < Min(len(words[i]), len(words[i-1])); j++ {
			if mp[words[i-1][j]] < mp[words[i][j]] {
				continue outer
			} else if mp[words[i-1][j]] > mp[words[i][j]] {
				return false
			}
		}
		if len(words[i-1]) > len(words[i]) {
			return false
		}
	}
	return true
}
