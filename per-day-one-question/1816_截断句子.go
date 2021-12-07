/*
-------------------------------------
# @Time    : 2021/12/6 15:31:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1816_截断句子.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "chopper is not a tanuki"
	k := 5
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	end := 0
	for i := 0; i < len(s) && k > 0; i++ {
		if s[i] == ' ' {
			k--
			end = i
		}
	}
	if k > 0 {
		return s
	}
	return s[:end]
}
