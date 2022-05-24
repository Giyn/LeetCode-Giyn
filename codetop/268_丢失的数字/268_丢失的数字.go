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

func missingNumber(nums []int) (ans int) {
	for i := 0; i <= len(nums); i++ {
		ans ^= i
	}
	for _, num := range nums {
		ans ^= num
	}
	return
}
