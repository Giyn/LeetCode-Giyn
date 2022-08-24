/*
-------------------------------------
# @Time    : 2022/8/25 0:10:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_067_把字符串转换成整数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	str := "4193 with words"
	fmt.Println(strToInt(str))
}

func strToInt(str string) int {
	abs := 0
	sign := 1
	i := 0
	// 去除前导空格
	for i < len(str) && str[i] == ' ' {
		i++
	}
	// 标记正负
	if i < len(str) {
		if str[i] == '-' {
			sign = -1
			i++
		} else if str[i] == '+' {
			i++
		}
	}
	for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		abs = 10*abs + int(str[i]-'0')
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return sign * abs
}
