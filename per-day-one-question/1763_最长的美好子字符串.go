/*
-------------------------------------
# @Time    : 2022/2/1 0:18:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1763_最长的美好子字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "YazaAay"
	fmt.Println(longestNiceSubstring(s))
}

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 97 && !strings.Contains(s, string(s[i]+32)) || (s[i] >= 97 && !strings.Contains(s, string(s[i]-32))) {
			s1, s2 := longestNiceSubstring(s[:i]), longestNiceSubstring(s[i+1:])
			if len(s1) >= len(s2) {
				return s1
			}
			return s2
		}
	}
	return s
}
