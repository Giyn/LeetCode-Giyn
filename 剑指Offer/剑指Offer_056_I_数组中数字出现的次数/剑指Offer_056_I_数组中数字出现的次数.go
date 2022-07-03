/*
-------------------------------------
# @Time    : 2022/7/3 20:07:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_056_I_数组中数字出现的次数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers(nums))
}

func singleNumbers(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum ^= num
	}
	diffIdx := 0
	for sum&1 == 0 {
		sum >>= 1
		diffIdx++
	}
	num1, num2 := 0, 0
	for _, num := range nums {
		if (num>>diffIdx)&1 == 1 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
