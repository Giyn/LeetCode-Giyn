/*
-------------------------------------
# @Time    : 2022/4/10 0:00:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 804_唯一摩尔斯密码词.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) (ans int) {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	mp := make(map[string]bool, len(words))
	for _, word := range words {
		trans := ""
		for _, ch := range word {
			trans += morse[ch-'a']
		}
		if !mp[trans] {
			ans += 1
			mp[trans] = true
		}
	}
	return
}
