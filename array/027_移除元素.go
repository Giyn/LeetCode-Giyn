/*
-------------------------------------
# @Time    : 2021/10/1 11:09:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 027_移除元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 5, 7}
	fmt.Println(removeElement(nums, 2))
}

func removeElement(nums []int, val int) int {
	// 用于覆盖待删除元素
	slowPointer := 0
	for fastPointer := 0; fastPointer < len(nums); fastPointer++ {
		if val != nums[fastPointer] {
			nums[slowPointer] = nums[fastPointer]
			slowPointer++
		}
	}
	return slowPointer
}
