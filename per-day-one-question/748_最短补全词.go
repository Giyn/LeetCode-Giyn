/*
-------------------------------------
# @Time    : 2021/12/10 14:08:45
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 748_最短补全词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	licensePlate := "OgEu755"
	words := []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}
	fmt.Println(shortestCompletingWord(licensePlate, words))
}

func shortestCompletingWord(licensePlate string, words []string) (ans string) {
	record := [26]int{}
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) {
			record[unicode.ToLower(ch)-'a']++
		}
	}
	for _, word := range words {
		tmp := record
		for _, ch := range word {
			if tmp[ch-'a'] > 0 {
				tmp[ch-'a']--
			}
			if tmp == [26]int{} && (len(word) < len(ans) || ans == "") {
				ans = word
				break
			}
		}
	}
	return
}
