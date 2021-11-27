/*
-------------------------------------
# @Time    : 2021/11/1 17:14:13
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
	if len(p) > len(s) {
		return
	}
	pRecord := [26]int{}
	for i := 0; i < len(p); i++ {
		pRecord[p[i]-'a']++
	}
	for i := 0; i <= len(s)-len(p); i++ {
		sRecord := [26]int{}
		for j, k := i, 0; k < len(p); k, j = k+1, j+1 {
			sRecord[s[j]-'a']++
		}
		if sRecord == pRecord {
			ans = append(ans, i)
		}
	}
	return
}
