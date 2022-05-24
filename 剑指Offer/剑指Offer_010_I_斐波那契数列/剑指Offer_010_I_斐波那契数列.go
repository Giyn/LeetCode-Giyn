/*
-------------------------------------
# @Time    : 2022/3/25 14:11:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_010_I_斐波那契数列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 5
	fmt.Println(fib(n))
}

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}
