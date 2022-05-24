/*
-------------------------------------
# @Time    : 2022/3/9 15:50:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_017_含有所有字符的最短字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) (ans string) {
	tm, sm := map[byte]int{}, map[byte]int{}
	var matchLen int
	for i := range t {
		tm[t[i]]++
	}
	for l, r := 0, 0; r < len(s); r++ {
		sm[s[r]]++
		if _, ok := tm[s[r]]; ok && sm[s[r]] == tm[s[r]] {
			matchLen++
		}
		for matchLen == len(tm) {
			if len(ans) == 0 || len(s[l:r+1]) < len(ans) {
				ans = s[l : r+1]
			}
			if _, ok := tm[s[l]]; ok && sm[s[l]] == tm[s[l]] {
				matchLen--
			}
			sm[s[l]]--
			l++
		}
	}
	return
}
