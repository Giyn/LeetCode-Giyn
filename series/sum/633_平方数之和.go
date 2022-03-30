/*
-------------------------------------
# @Time    : 2022/3/30 16:55:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 633_平方数之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	c := 5
	fmt.Println(judgeSquareSum(c))
}

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		if left*left+right*right < c {
			left++
		} else if left*left+right*right > c {
			right--
		} else {
			return true
		}
	}
	return false
}
