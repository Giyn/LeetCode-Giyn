/*
-------------------------------------
# @Time    : 2022/4/6 20:31:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 009_回文数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	x := 1234321
	fmt.Println(isPalindromeNum(x))
}

func isPalindromeNum(x int) bool {
	// 没有前导0
	if (x < 0) || (x != 0 && x%10 == 0) {
		return false
	}
	// 反转右半部分
	reversed := 0
	for x > reversed {
		reversed = 10*reversed + x%10
		x /= 10
	}
	return x == reversed || x == reversed/10
}
