/*
-------------------------------------
# @Time    : 2022/8/29 23:48:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 260_只出现一次的数字 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
