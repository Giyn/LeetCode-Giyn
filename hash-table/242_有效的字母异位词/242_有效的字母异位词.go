/*
-------------------------------------
# @Time    : 2021/11/1 9:21:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 242_有效的字母异位词.go
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
	if len(s) != len(t) {
		return false
	}
	mp := make(map[rune]int)
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

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	record := make(map[byte]int)
//	for i := 0; i < len(s); i++ {
//		record[s[i]-'a']++
//	}
//	for i := 0; i < len(t); i++ {
//		if record[t[i]-'a'] > 0 {
//			record[t[i]-'a']--
//		} else {
//			return false
//		}
//	}
//	return true
//}
