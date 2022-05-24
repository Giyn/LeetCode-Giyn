/*
-------------------------------------
# @Time    : 2022/3/9 10:53:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_014_字符串中的变位词.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	mp1 := [26]int{}
	mp2 := [26]int{}
	for i, ch := range s1 {
		mp1[ch-'a']++
		mp2[s2[i]-'a']++
	}
	if mp1 == mp2 {
		return true
	}
	for right := len(s1); right < len(s2); right++ {
		mp2[s2[right]-'a']++
		mp2[s2[right-len(s1)]-'a']--
		if mp1 == mp2 {
			return true
		}
	}
	return false
}
