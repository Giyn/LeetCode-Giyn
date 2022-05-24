/*
-------------------------------------
# @Time    : 2022/5/8 1:59:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 442_数组中重复的数据.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) (ans []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			ans = append(ans, num)
		}
	}
	return
}
