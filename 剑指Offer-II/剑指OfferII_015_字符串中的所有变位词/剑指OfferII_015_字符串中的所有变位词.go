/*
-------------------------------------
# @Time    : 2022/3/9 11:55:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_015_字符串中的所有变位词.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) (ans []int) {
	if len(s) < len(p) {
		return
	}
	mp1 := [26]int{}
	mp2 := [26]int{}
	for i, ch := range p {
		mp1[ch-'a']++
		mp2[s[i]-'a']++
	}
	if mp1 == mp2 {
		ans = append(ans, 0)
	}
	for right := len(p); right < len(s); right++ {
		mp2[s[right]-'a']++
		mp2[s[right-len(p)]-'a']--
		if mp1 == mp2 {
			ans = append(ans, right-len(p)+1)
		}
	}
	return
}
