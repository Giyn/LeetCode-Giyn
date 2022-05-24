/*
-------------------------------------
# @Time    : 2022/3/4 1:12:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_004_只出现一次的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 100}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		bitNum := 0
		for _, num := range nums {
			bitNum += (num >> i) & 1
		}
		if bitNum%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
