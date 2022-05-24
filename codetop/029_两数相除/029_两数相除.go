/*
-------------------------------------
# @Time    : 2022/4/28 12:02:31
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
	dividend := 10
	divisor := 3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	diffSign := false
	if (dividend < 0) != (divisor < 0) {
		diffSign = true
	}
	a, b := abs(dividend), abs(divisor)
	ans := 0
	for a >= b {
		base, cnt := b, 1
		for a >= base<<1 {
			base <<= 1
			cnt <<= 1
		}
		a -= base
		ans += cnt
	}
	if diffSign {
		return -ans
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
