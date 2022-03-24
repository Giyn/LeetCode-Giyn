/*
-------------------------------------
# @Time    : 2022/3/24 9:37:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 661_图片平滑器.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	img := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	fmt.Println(imageSmoother(img))
}

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			sum, count := 0, 0
			for _, row := range img[Max(i-1, 0):Min(i+2, m)] {
				for _, v := range row[Max(j-1, 0):Min(j+2, n)] {
					sum += v
					count++
				}
			}
			ans[i][j] = sum / count
		}
	}
	return ans
}
