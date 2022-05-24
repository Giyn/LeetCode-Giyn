/*
-------------------------------------
# @Time    : 2022/2/4 16:48:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1725_可以形成最大正方形的矩形数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	fmt.Println(countGoodRectangles(rectangles))
}

func countGoodRectangles(rectangles [][]int) (ans int) {
	maxLen := -1
	for _, rectangle := range rectangles {
		tmp := min(rectangle[0], rectangle[1])
		if tmp > maxLen {
			maxLen = tmp
			ans = 1
		} else if tmp == maxLen {
			ans++
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
