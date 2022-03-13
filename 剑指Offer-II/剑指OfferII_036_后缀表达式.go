/*
-------------------------------------
# @Time    : 2022/3/12 19:29:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_036_后缀表达式.go
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
	for _, token := range tokens {
		if v, err := strconv.Atoi(token); err == nil {
			stack = append(stack, v)
		} else {
			if token == "+" {
				res := stack[len(stack)-2] + stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			} else if token == "-" {
				res := stack[len(stack)-2] - stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, res)
			} else if token == "*" {
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
