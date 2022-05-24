/*
-------------------------------------
# @Time    : 2022/5/19 13:32:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_020_表示数值的字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "-."
	fmt.Println(isNumber(s))
}

func isNumber(s string) bool {
	isDecimal := func(s string) bool {
		if strings.Count(s, ".") != 1 || len(s) == 0 {
			return false
		}
		left, right := 0, 0
		start := 0
		if s[0] == '+' || s[0] == '-' {
			if len(s) == 1 {
				return false
			}
			start++
		}
		idx := strings.Index(s, ".")
		for _, ch := range s[start:idx] {
			if !unicode.IsDigit(ch) {
				return false
			}
			left++
		}
		for _, ch := range s[idx+1:] {
			if !unicode.IsDigit(ch) {
				return false
			}
			right++
		}
		return left >= 1 || (left == 0 && right > 0)
	}
	isInteger := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		start := 0
		if s[0] == '+' || s[0] == '-' {
			if len(s) == 1 {
				return false
			}
			start++
		}
		for _, ch := range s[start:] {
			if !unicode.IsDigit(ch) {
				return false
			}
		}
		return true
	}
	s = strings.TrimSpace(s)
	if strings.Count(s, "e")+strings.Count(s, "E") > 1 {
		return false
	}
	idx := max(-1, max(strings.Index(s, "e"), strings.Index(s, "E")))
	if idx == -1 {
		return isDecimal(s) || isInteger(s)
	} else {
		return (isDecimal(s[:idx]) || isInteger(s[:idx])) && isInteger(s[idx+1:])
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
