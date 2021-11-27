/*
-------------------------------------
# @Time    : 2021/11/1 10:00:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 383_赎金信.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	ransomNote := "aa"
	magazine := "aba"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	mp := make(map[rune]int)
	for _, ch := range magazine {
		mp[ch]++
	}

	for _, ch := range ransomNote {
		if v, ok := mp[ch]; v > 0 && ok {
			mp[ch]--
		} else {
			return false
		}
	}
	return true
}
