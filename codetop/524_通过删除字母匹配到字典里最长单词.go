/*
-------------------------------------
# @Time    : 2022/5/2 15:49:53
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
	s := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}
	fmt.Println(findLongestWord(s, dictionary))
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return len(dictionary[i]) > len(dictionary[j])
	})
	for _, word := range dictionary {
		i, j := 0, 0
		for i < len(s) {
			if s[i] == word[j] {
				j++
				if j == len(word) {
					return word
				}
			}
			i++
		}
	}
	return ""
}
