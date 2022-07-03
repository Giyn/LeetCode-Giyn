/*
-------------------------------------
# @Time    : 2022/7/3 20:41:45
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_056_II_数组中数字出现的次数 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) (ans int) {
	for i := 0; i < 32; i++ {
		bit := 0
		for _, num := range nums {
			bit += (num >> i) & 1
		}
		ans += (bit % 3) << i
	}
	return
}
