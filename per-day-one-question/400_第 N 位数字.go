/*
-------------------------------------
# @Time    : 2021/11/30 8:54:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 400_第 N 位数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 11
	fmt.Println(findNthDigit(n))
}

func findNthDigit(n int) int {
	// 判断是几位数
	digitBit := 1
	base := 9
	for n > digitBit*base {
		n -= digitBit * base
		digitBit++
		base *= 10
	}
	// 从0开始
	n--
	// 求出该数字
	num := int(math.Pow10(digitBit-1)) + n/digitBit
	// 求出目标值在该数字的索引
	index := n % digitBit
	return num / int(math.Pow10(digitBit-1-index)) % 10
}
