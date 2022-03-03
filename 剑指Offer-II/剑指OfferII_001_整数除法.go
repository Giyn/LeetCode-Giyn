/*
-------------------------------------
# @Time    : 2022/3/1 23:45:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_001_整数除法.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"LeetCodeGiyn/utils"
	"fmt"
	"math"
)

func main() {
	a := 7
	b := -3
	fmt.Println(divide(a, b))
}

func divide(a int, b int) (ans int) {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	if a == 0 || b == 1 {
		return a
	} else if b == -1 {
		return -a
	} else if a == b {
		return 1
	}

	sign := 1
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		sign = -1
	}
	a, b = utils.Abs(a), utils.Abs(b)
	calc := func(x, y int) (int, int) {
		n := 1
		for x > y<<1 {
			y <<= 1
			n <<= 1
		}
		return n, y
	}
	for a >= b {
		cnt, val := calc(a, b)
		ans += cnt
		a -= val
	}
	return sign * ans
}
