/*
-------------------------------------
# @Time    : 2022/3/7 0:11:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 504_七进制数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := -7
	fmt.Println(convertToBase7(num))
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var s []string
	var neg bool
	if num < 0 {
		neg = true
		num = -num
	}
	for num > 0 {
		s = append(s, strconv.Itoa(num%7))
		num /= 7
	}
	if neg {
		s = append(s, "-")
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return strings.Join(s, "")
}
