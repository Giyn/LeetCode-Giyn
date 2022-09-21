/*
-------------------------------------
# @Time    : 2022/9/21 23:18:22
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_044_数字序列中某一位的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 200
	fmt.Println(findNthDigit(n))
}

func findNthDigit(n int) int {
	bit := 1
	base := 9
	// 1: 9
	// 2: 90
	// 3: 900
	// ...
	for n > bit*base {
		n -= bit * base
		bit++
		base *= 10
	}
	n--
	num := int(math.Pow10(bit-1)) + n/bit
	idx := n % bit
	return num / int(math.Pow10(bit-1-idx)) % 10
}
