/*
-------------------------------------
# @Time    : 2022/1/30 0:14:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 884_两句话中的不常见单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println(uncommonFromSentences(s1, s2))
}

func uncommonFromSentences(s1 string, s2 string) (ans []string) {
	s1Slice := strings.Fields(s1)
	s2Slice := strings.Fields(s2)
	s1Map := make(map[string]int, len(s1Slice))
	s2Map := make(map[string]int, len(s2Slice))
	for _, word := range s1Slice {
		s1Map[word]++
	}
	for _, word := range s2Slice {
		s2Map[word]++
	}
	for _, word := range s2Slice {
		if _, ok := s1Map[word]; !ok && s2Map[word] == 1 {
			ans = append(ans, word)
		}
	}
	for _, word := range s1Slice {
		if _, ok := s2Map[word]; !ok && s1Map[word] == 1 {
			ans = append(ans, word)
		}
	}
	return
}
