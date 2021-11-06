/*
-------------------------------------
# @Time    : 2021/11/6 12:11:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 268_丢失的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return (len(nums)+1)*len(nums)/2 - sum
}
