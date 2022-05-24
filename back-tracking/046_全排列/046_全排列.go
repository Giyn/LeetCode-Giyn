/*
-------------------------------------
# @Time    : 2021/12/7 15:47:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 046_全排列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) (ans [][]int) {
	var path []int
	var search func([]int, int) bool
	search = func(s []int, val int) bool {
		for _, v := range s {
			if val == v {
				return true
			}
		}
		return false
	}
	var backtrack func(int)
	backtrack = func(index int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if search(path, nums[i]) {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
