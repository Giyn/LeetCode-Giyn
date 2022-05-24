/*
-------------------------------------
# @Time    : 2022/5/22 15:36:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 402_移掉 K 位数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	num := "10200"
	k := 1
	fmt.Println(removeKdigits(num, k))
}

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		return "0"
	}
	return ans
}
