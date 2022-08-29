/*
-------------------------------------
# @Time    : 2022/8/29 23:36:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 136_只出现一次的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	ans := nums[0]
	for _, num := range nums[1:] {
		ans ^= num
	}
	return ans
}
