/*
-------------------------------------
# @Time    : 2022/4/3 13:57:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 678_有效的括号字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "(*))"
	fmt.Println(checkValidString(s))
}

func checkValidString(s string) bool {
	minLeft, maxLeft := 0, 0
	for _, ch := range s {
		if ch == '(' {
			minLeft++
			maxLeft++
		} else if ch == ')' {
			minLeft = max(0, minLeft-1)
			maxLeft--
			if maxLeft < 0 {
				return false
			}
		} else {
			minLeft = max(0, minLeft-1)
			maxLeft++
		}
	}
	return minLeft == 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
