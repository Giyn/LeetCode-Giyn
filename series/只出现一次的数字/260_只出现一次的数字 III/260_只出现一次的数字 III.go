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

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 1, 2, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	xorSum := 0
	for _, num := range nums {
		xorSum ^= num
	}
	var lsb int
	if xorSum == math.MinInt64 {
		lsb = xorSum // 负0的特点是第一位是1，其余位是0，所以它的最低有效位就是自己
	} else {
		lsb = xorSum & -xorSum // 最低有效位
	}
	//k := -1
	//for i := 31; i >= 0 && k == -1; i-- {
	//	if ((xorSum >> i) & 1) == 1 {
	//		k = i
	//	}
	//}
	type1, type2 := 0, 0
	for _, num := range nums {
		//if ((num >> k) & 1) == 1 {
		//	type1 ^= num
		//} else {
		//	type2 ^= num
		//}
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	return []int{type1, type2}
}
