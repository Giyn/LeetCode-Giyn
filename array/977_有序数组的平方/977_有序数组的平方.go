/*
-------------------------------------
# @Time    : 2021/10/13 10:33:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 977_有序数组的平方.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j := 0, n-1
	ans := make([]int, n)

	for pos := n - 1; pos >= 0; pos-- {
		if a, b := nums[i]*nums[i], nums[j]*nums[j]; a > b {
			ans[pos] = a
			i++
		} else {
			ans[pos] = b
			j--
		}
	}
	return ans
}

//func sortedSquares(nums []int) []int {
//	for i, v := range nums {
//		nums[i] = v * v
//	}
//	sort.Ints(nums)
//
//	return nums
//}
