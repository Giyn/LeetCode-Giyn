/*
-------------------------------------
# @Time    : 2022/4/2 23:55:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 415_字符串相加.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "1"
	num2 := "9"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) (ans string) {
	l1, l2 := len(num1)-1, len(num2)-1
	carry := 0
	for l1 >= 0 || l2 >= 0 || carry != 0 {
		var x, y int
		if l1 >= 0 {
			x = int(num1[l1] - '0')
		}
		if l2 >= 0 {
			y = int(num2[l2] - '0')
		}
		sum := x + y + carry
		ans = strconv.Itoa(sum%10) + ans
		carry = sum / 10
		l1--
		l2--
	}
	return
}
