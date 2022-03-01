/*
-------------------------------------
# @Time    : 2022/3/1 19:09:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 011_盛最多水的容器.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = utils.Max(ans, (right-left)*utils.Min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}
