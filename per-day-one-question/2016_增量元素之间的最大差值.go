/*
-------------------------------------
# @Time    : 2022/2/26 10:08:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2016_增量元素之间的最大差值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 4}
	fmt.Println(maximumDifference(nums))
}

func maximumDifference(nums []int) (ans int) {
	min := nums[0]
	ans = -1
	for _, num := range nums {
		diff := num - min
		if diff > 0 && diff > ans {
			ans = diff
		} else if diff < 0 {
			min = num
		}
	}
	return
}
