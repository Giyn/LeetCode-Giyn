/*
-------------------------------------
# @Time    : 2022/4/7 13:01:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 227_基本计算器 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := " 3+5 / 2 "
	fmt.Println(calculate(s))
}

func calculate(s string) (ans int) {
	var stack []int
	var num int
	sign := '+'
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'
		if isDigit {
			num = 10*num + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			sign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}
