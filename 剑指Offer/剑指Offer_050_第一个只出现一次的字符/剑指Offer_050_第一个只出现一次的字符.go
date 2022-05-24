/*
-------------------------------------
# @Time    : 2022/5/17 4:29:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_050_第一个只出现一次的字符.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abaccdeff"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) byte {
	mp := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if v, ok := mp[s[i]]; ok && v == 1 {
			return s[i]
		}
	}
	return ' '
}
