/*
-------------------------------------
# @Time    : 2022/5/19 19:58:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_031_栈的压入、弹出序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	j := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}
