/*
-------------------------------------
# @Time    : 2022/7/10 23:00:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_064_求1+2+…+n.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 100
	fmt.Println(sumNums(n))
}

func sumNums(n int) int {
	_ = n > 0 && func() bool { n += sumNums(n - 1); return true }()
	return n
}
