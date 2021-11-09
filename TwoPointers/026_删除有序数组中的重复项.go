/*
-------------------------------------
# @Time    : 2021/10/1 11:09:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 026_删除有序数组中的重复项.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 3, 3, 5, 6, 7}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slowPointer := 0
	for fastPointer := 1; fastPointer < len(nums); fastPointer++ {
		if nums[slowPointer] != nums[fastPointer] {
			slowPointer++
			nums[slowPointer] = nums[fastPointer]
		}
	}
	return slowPointer + 1
}
