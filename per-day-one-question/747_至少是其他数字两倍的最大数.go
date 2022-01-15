/*
-------------------------------------
# @Time    : 2022/1/14 19:04:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 747_至少是其他数字两倍的最大数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(nums))
}

func dominantIndex(nums []int) int {
	max, second, index := -1, -1, -1
	for i, num := range nums {
		if num > max {
			second = max
			max = num
			index = i
		} else if num > second {
			second = num
		}
	}
	if second*2 <= max {
		return index
	} else {
		return -1
	}
}
