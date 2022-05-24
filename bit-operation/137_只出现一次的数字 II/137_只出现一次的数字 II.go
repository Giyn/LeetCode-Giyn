/*
-------------------------------------
# @Time    : 2022/3/4 10:05:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 137_只出现一次的数字 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
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
