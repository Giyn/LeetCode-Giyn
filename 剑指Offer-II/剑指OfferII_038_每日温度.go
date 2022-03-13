/*
-------------------------------------
# @Time    : 2022/3/12 21:49:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_038_每日温度.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) (ans []int) {
	n := len(temperatures)
	var stack []int
	ans = make([]int, n)
	for i, temperature := range temperatures {
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return
}
