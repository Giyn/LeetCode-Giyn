/*
-------------------------------------
# @Time    : 2021/10/12 8:53:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 029_两数相除.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	dividend, divisor := 10, 3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	abs := func(n int) int {
		y := n >> 63
		return (n ^ y) - y
	}
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
	diffSign := false
	if (dividend < 0) != (divisor < 0) {
		diffSign = true
	}
	a, b := abs(dividend), abs(divisor)
	ans := 0
	for a >= b {
		base, cnt := b, 1
		for (base << 1) <= a {
			base <<= 1
			cnt <<= 1
		}
		ans += cnt
		a -= base
	}
	if diffSign {
		ans = -ans
	}

	return ans
}
