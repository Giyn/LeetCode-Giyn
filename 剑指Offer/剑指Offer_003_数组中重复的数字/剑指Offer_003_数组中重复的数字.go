/*
-------------------------------------
# @Time    : 2022/3/22 9:39:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_003_数组中重复的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	mp := make([]int, len(nums))
	for _, num := range nums {
		mp[num]++
		if mp[num] > 1 {
			return num
		}
	}
	return 0
}
