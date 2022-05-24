/*
-------------------------------------
# @Time    : 2022/5/17 0:03:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 953_验证外星语词典.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}

func isAlienSorted(words []string, order string) bool {
	mp := make(map[byte]int, 26)
	for i, ch := range order {
		mp[byte(ch)] = i
	}
	for i := 1; i < len(words); i++ {
		if strings.HasPrefix(words[i-1], words[i]) && words[i-1] != words[i] {
			return false
		}
		for j := 0; j < len(words[i-1]) && j < len(words[i]); j++ {
			if mp[words[i-1][j]] > mp[words[i][j]] {
				return false
			} else if mp[words[i-1][j]] < mp[words[i][j]] {
				break
			}
		}
	}
	return true
}
