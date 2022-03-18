/*
-------------------------------------
# @Time    : 2021/11/10 1:19:22
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 150_逆波兰表达式求值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	var stack []int
	for _, ch := range tokens {
		if i, err := strconv.Atoi(ch); err == nil {
			stack = append(stack, i)
		} else {
			if ch == "+" {
				res := stack[len(stack)-2] + stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			} else if ch == "-" {
				res := stack[len(stack)-2] - stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			} else if ch == "*" {
				res := stack[len(stack)-2] * stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			} else {
				res := stack[len(stack)-2] / stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			}
		}
	}
	return stack[len(stack)-1]
}
