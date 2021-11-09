/*
-------------------------------------
# @Time    : 2021/10/1 13:58:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 283_移动零.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	slowPointer := 0
	for fastPointer := 0; fastPointer < len(nums); fastPointer++ {
		if nums[fastPointer] != 0 {
			nums[slowPointer], nums[fastPointer] = nums[fastPointer], nums[slowPointer]
			slowPointer++
		}
	}
}
