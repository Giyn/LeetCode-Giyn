/*
-------------------------------------
# @Time    : 2022/4/28 17:28:12
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

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}
	for i, num := range nums {
		if v, ok := mp[target-num]; ok {
			return []int{i, v}
		} else {
			mp[num] = i
		}
	}
	return []int{}
}
