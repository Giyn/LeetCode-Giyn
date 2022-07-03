/*
-------------------------------------
# @Time    : 2022/7/2 18:54:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_053_II_0～n-1中缺失的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
