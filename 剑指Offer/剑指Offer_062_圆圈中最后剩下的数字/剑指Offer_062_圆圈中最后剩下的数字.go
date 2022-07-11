/*
-------------------------------------
# @Time    : 2022/7/10 22:05:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_062_圆圈中最后剩下的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 10
	m := 17
	fmt.Println(lastRemaining(n, m))
}

func lastRemaining(n int, m int) int {
	idx := 0
	for i := 2; i <= n; i++ {
		idx = (idx + m) % i
	}
	return idx
}
