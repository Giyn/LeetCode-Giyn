/*
-------------------------------------
# @Time    : 2022/3/29 21:29:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 008_字符串转换整数 (atoi).go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	s := "4193 with words"
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	abs, sign, i := 0, 1, 0
	// 去除前导空格
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// 标记正负
	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0')
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}
	return sign * abs
}
