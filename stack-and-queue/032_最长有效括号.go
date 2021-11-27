/*
-------------------------------------
# @Time    : 2021/11/10 16:57:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 032_最长有效括号.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "(()))())("
	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	max := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left < right {
			left, right = 0, 0
		} else if left == right {
			max = max032(max, 2*left)
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			left, right = 0, 0
		} else if left == right {
			max = max032(max, 2*right)
		}
	}
	return max
}

//func longestValidParentheses(s string) int {
//	stack := []int{-1}
//	max := 0
//	for i, ch := range s {
//		if ch == '(' {
//			stack = append(stack, i)
//		} else {
//			stack = stack[:len(stack)-1]
//			if len(stack) == 0 {
//				stack = append(stack, i)
//			}
//			max = max032(max, i-stack[len(stack)-1])
//		}
//	}
//	return max
//}

func max032(x, y int) int {
	if x > y {
		return x
	}
	return y
}
