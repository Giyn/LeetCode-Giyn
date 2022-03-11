/*
-------------------------------------
# @Time    : 2022/3/4 9:52:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 029_两数相除.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
	"math"
)

func main() {
	dividend := 10
	divisor := 3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) (ans int) {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}
	if dividend == math.MinInt32 && divisor == math.MinInt32 {
		return 1
	} else if divisor == math.MinInt32 {
		return 0
	}
	diffSign := 1
	if (dividend < 0) != (divisor < 0) {
		diffSign = -1
	}
	a, b := Abs(dividend), Abs(divisor)
	for a >= b {
		base, cnt := b, 1
		for (base << 1) <= a {
			base <<= 1
			cnt <<= 1
		}
		ans += cnt
		a -= base
	}
	return diffSign * ans
}
