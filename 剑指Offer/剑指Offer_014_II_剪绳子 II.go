/*
-------------------------------------
# @Time    : 2022/3/31 20:18:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_014_II_剪绳子 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	n := 10
	fmt.Println(cuttingRope2(n))
}

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		ans = ans * 3 % (1e9 + 7)
		n -= 3
	}
	return (ans * n) % (1e9 + 7)
}
