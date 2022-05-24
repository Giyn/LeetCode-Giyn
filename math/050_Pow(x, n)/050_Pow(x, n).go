/*
-------------------------------------
# @Time    : 2021/12/5 16:48:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 050_Pow(x, n).go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

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
	for n != 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}
