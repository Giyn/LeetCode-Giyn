/*
-------------------------------------
# @Time    : 2021/12/2 8:43:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 524_通过删除字母匹配到字典里最长单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "aewfafwafjlwajflwajflwafj"
	dictionary := []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}
	fmt.Println(findLongestWord(s, dictionary))
}

func findLongestWord(s string, dictionary []string) (ans string) {
	sort.Slice(dictionary, func(i, j int) bool {
		a, b := dictionary[i], dictionary[j]
		return len(a) > len(b) || len(a) == len(b) && a < b
	})
	for _, word := range dictionary {
		i := 0
		for j := range s {
			if s[j] == word[i] {
				i++
				if i == len(word) {
					return word
				}
			}
		}
	}
	return
}
