/*
-------------------------------------
# @Time    : 2021/11/1 23:59:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 001_两数之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) (ans []int) {
	mp := make(map[int]int)
	for index, num := range nums {
		if preIndex, ok := mp[target-num]; ok {
			return []int{index, preIndex}
		} else {
			mp[num] = index
		}
	}
	return
}
