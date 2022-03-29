/*
-------------------------------------
# @Time    : 2022/3/29 23:43:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 287_寻找重复数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) int {
	left, right := 1, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
