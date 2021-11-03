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

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

func trap(height []int) (ans int) {
	var stack []int
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			right := i
			h := min042(height[left], height[right]) - height[cur]
			ans += (right - left - 1) * h
		}
		stack = append(stack, i)
	}
	return
}

func min042(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// 暴力解法
//func trap(height []int) (ans int) {
//	for i := 1; i < len(height)-1; i++ {
//		maxLeft, maxRight := 0, 0
//		for j := i; j >= 0; j-- {
//			maxLeft = max042(maxLeft, height[j])
//		}
//		for j := i; j < len(height); j++ {
//			maxRight = max042(maxRight, height[j])
//		}
//		ans += min042(maxLeft, maxRight) - height[i]
//	}
//	return
//}

//func max042(x, y int) int {
//	if x > y {
//		return x
//	} else {
//		return y
//	}
//}
