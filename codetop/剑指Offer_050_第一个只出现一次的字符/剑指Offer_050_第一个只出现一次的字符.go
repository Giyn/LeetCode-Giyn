/*
-------------------------------------
# @Time    : 2022/6/27 2:45:42
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
	n := len(s)
	mp := make(map[byte]int, n)
	for i := 0; i < n; i++ {
		mp[s[i]]++
	}
	for i := 0; i < n; i++ {
		if v := mp[s[i]]; v == 1 {
			return s[i]
		}
	}
	return ' '
}
