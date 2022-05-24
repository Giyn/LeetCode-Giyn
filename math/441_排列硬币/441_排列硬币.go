/*
-------------------------------------
# @Time    : 2021/12/5 16:51:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 441_排列硬币.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 8
	fmt.Println(arrangeCoins(n))
}

func arrangeCoins(n int) int {
	return int(math.Floor((math.Sqrt(float64(1+8*n)) - 1) / 2))
}
