/*
-------------------------------------
# @Time    : 2022/4/17 0:03:51
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 819_最常见的单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}

func mostCommonWord(paragraph string, banned []string) string {
	ban := map[string]bool{}
	for _, s := range banned {
		ban[s] = true
	}
	freq := map[string]int{}
	maxFreq := 0
	var word []byte
	for i := 0; i <= len(paragraph); i++ {
		if i < len(paragraph) && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil {
			s := string(word)
			if !ban[s] {
				freq[s]++
				maxFreq = max(maxFreq, freq[s])
			}
			word = nil
		}
	}
	for s, f := range freq {
		if f == maxFreq {
			return s
		}
	}
	return ""
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
