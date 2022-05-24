/*
-------------------------------------
# @Time    : 2022/3/12 17:01:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_032_有效的变位词.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	mp := map[rune]int{}
	for _, ch := range s {
		mp[ch]++
	}
	for _, ch := range t {
		if mp[ch] > 0 {
			mp[ch]--
		} else {
			return false
		}
	}
	return true
}
