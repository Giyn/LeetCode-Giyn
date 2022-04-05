/*
-------------------------------------
# @Time    : 2022/4/5 19:11:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_065_不用加减乘除做加法.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	a := 1
	b := 1
	fmt.Println(add(a, b))
}

func add(a int, b int) int {
	carry := 0
	for b != 0 {
		carry = (a & b) << 1 // 进位
		a ^= b               // 无进位和
		b = carry
	}
	return a
}
