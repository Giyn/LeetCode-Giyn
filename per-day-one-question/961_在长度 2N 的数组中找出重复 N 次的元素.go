/*
-------------------------------------
# @Time    : 2022/5/21 2:45:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 961_在长度 2N 的数组中找出重复 N 次的元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Println(repeatedNTimes(nums))
}

func repeatedNTimes(nums []int) int {
	mp := make(map[int]bool)
	for _, num := range nums {
		if mp[num] {
			return num
		}
		mp[num] = true
	}
	return 0
}
