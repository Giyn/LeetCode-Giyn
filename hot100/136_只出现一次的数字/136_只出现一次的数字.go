/*
-------------------------------------
# @Time    : 2022/2/26 23:10:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 136_只出现一次的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) (ans int) {
	for _, num := range nums {
		ans ^= num
	}
	return
}
