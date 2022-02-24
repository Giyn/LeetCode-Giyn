/*
-------------------------------------
# @Time    : 2022/2/24 11:17:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 739_每日温度.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var stack []int
	for i, v := range temperatures {
		for len(stack) > 0 && v > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}
