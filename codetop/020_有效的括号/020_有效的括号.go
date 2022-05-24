/*
-------------------------------------
# @Time    : 2021/11/9 22:52:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 020_有效的括号.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			if (stack[len(stack)-1] == '(' && s[i] == ')') || (stack[len(stack)-1] == '{' && s[i] == '}') || (stack[len(stack)-1] == '[' && s[i] == ']') {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
