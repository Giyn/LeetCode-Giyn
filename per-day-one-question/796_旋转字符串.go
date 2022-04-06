/*
-------------------------------------
# @Time    : 2022/4/7 1:36:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 796_旋转字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcde"
	goal := "cdeab"
	fmt.Println(rotateString(s, goal))
}

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
