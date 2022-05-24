/*
-------------------------------------
# @Time    : 2021/12/4 9:18:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 383_赎金信.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	record := [26]int{}
	for i := 0; i < len(magazine); i++ {
		record[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if record[ransomNote[i]-'a'] > 0 {
			record[ransomNote[i]-'a']--
		} else {
			return false
		}
	}
	return true
}
