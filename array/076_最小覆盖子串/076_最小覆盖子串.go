/*
-------------------------------------
# @Time    : 2021/10/14 22:18:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 076_最小覆盖子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s, t := "ADOBECODEBANC", "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	tm, sm := map[byte]int{}, map[byte]int{}
	var matchLen int
	var ans string

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
	return ans
}
