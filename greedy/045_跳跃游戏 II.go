/*
-------------------------------------
# @Time    : 2021/12/14 14:58:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 045_跳跃游戏 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(jump(nums))
}

func jump(nums []int) (ans int) {
	var curCover int
	var nextCover int
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > nextCover {
			nextCover = i + nums[i]
		}
		if i == curCover {
			curCover = nextCover
			ans++
		}
	}
	return
}
