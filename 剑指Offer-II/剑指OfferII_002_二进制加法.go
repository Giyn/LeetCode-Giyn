/*
-------------------------------------
# @Time    : 2022/3/2 1:20:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_002_二进制加法.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	var ans, carry int
	for y > 0 {
		ans = x ^ y          // 无进位结果
		carry = (x & y) << 1 // 进位
		x, y = ans, carry
	}
	return strconv.Itoa(ans)
}
