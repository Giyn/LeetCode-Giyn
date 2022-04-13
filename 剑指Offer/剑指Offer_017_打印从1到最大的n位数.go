/*
-------------------------------------
# @Time    : 2022/4/10 17:16:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_017_打印从1到最大的n位数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 3
	fmt.Println(printNumbers(n))
}

func printNumbers(n int) (ans []int) {
	bound := 10
	for n > 1 {
		bound *= 10
		n--
	}
	for i := 1; i < bound; i++ {
		ans = append(ans, i)
	}
	return
}
