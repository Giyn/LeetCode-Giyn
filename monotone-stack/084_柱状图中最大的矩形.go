/*
-------------------------------------
# @Time    : 2022/2/26 10:18:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 084_柱状图中最大的矩形.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) (ans int) {
	var stack []int
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack = append(stack, 0)
	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if heights[i] == heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					left := stack[len(stack)-1]
					right := i
					w := right - left - 1
					h := heights[mid]
					ans = utils.Max(ans, w*h)
				}
			}
			stack = append(stack, i)
		}
	}
	return
}
