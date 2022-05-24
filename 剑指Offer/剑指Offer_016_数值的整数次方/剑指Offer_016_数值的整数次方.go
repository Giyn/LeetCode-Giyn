/*
-------------------------------------
# @Time    : 2022/5/19 12:05:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_016_数值的整数次方.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	x := 0.00001
	n := 2147483647
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ans := 1.0
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}
