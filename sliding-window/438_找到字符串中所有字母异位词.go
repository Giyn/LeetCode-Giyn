/*
-------------------------------------
# @Time    : 2022/3/9 14:43:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 438_找到字符串中所有字母异位词.go
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
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}
	sRecord, pRecord := [26]int{}, [26]int{}
	for i, ch := range p {
		sRecord[s[i]-'a']++
		pRecord[ch-'a']++
	}
	if sRecord == pRecord {
		ans = append(ans, 0)
	}
	for i, ch := range s[:sLen-pLen] {
		sRecord[ch-'a']--
		sRecord[s[i+pLen]-'a']++
		if sRecord == pRecord {
			ans = append(ans, i+1)
		}
	}
	return
}
