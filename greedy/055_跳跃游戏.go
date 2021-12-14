/*
-------------------------------------
# @Time    : 2021/12/14 8:47:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 055_跳跃游戏.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	var cover int
	//for i := 0; i <= cover; i++ {
	//	if i+nums[i] > cover {
	//		cover = i + nums[i]
	//	}
	//	if cover >= len(nums)-1 {
	//		return true
	//	}
	//}
	//return false
	for i := 0; i < len(nums); i++ {
		if i > cover {
			return false
		}
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
	}
	return true
}
