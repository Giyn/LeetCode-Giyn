/*
-------------------------------------
# @Time    : 2021/11/3 0:05:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 042_接雨水.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

func trap(height []int) (ans int) {
	// 单调栈
	var stack []int
	stack = append(stack, 0)
	for i := 1; i < len(height); i++ {
		if height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					left := stack[len(stack)-1]
					right := i
					h := min(height[left], height[right]) - height[mid]
					ans += (right - left - 1) * h
				}
			}
			stack = append(stack, i)
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
