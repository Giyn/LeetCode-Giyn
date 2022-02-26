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
	"LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//height := []int{4, 2, 0, 3, 2, 5}
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
					h := utils.Min(height[left], height[right]) - height[mid]
					ans += (right - left - 1) * h
				}
			}
			stack = append(stack, i)
		}
	}
	return

	// 动态规划
	//n := len(height)
	//if n < 3 {
	//	return
	//}
	//maxLeft := make([]int, n)
	//maxRight := make([]int, n)
	//maxLeft[0] = height[0]
	//maxRight[n-1] = height[n-1]
	//for i := 1; i < n; i++ {
	//	maxLeft[i] = utils.Max(maxLeft[i-1], height[i])
	//}
	//for i := n - 2; i >= 0; i-- {
	//	maxRight[i] = utils.Max(maxRight[i+1], height[i])
	//}
	//for i := 0; i < n; i++ {
	//	ans += utils.Min(maxLeft[i], maxRight[i]) - height[i]
	//}
	//return
}
