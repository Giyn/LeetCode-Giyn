/*
-------------------------------------
# @Time    : 2021/10/30 22:23:11
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
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	k := 0
	for i := sum; i&1 == 0; i >>= 1 {
		k++
	}
	ans := []int{0, 0}
	for _, v := range nums {
		if (v>>k)&1 == 1 {
			ans[1] ^= v
		} else {
			ans[0] ^= v
		}
	}
	return ans
}
