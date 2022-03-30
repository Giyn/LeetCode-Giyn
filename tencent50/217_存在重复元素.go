/*
-------------------------------------
# @Time    : 2022/3/30 20:39:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 217_存在重复元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool, len(nums))
	for _, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true
	}
	return false
}
