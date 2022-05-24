/*
-------------------------------------
# @Time    : 2022/3/31 23:02:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_015_二进制中1的个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	var n uint32 = 128
	fmt.Println(hammingWeight(n))
}

func hammingWeight(num uint32) (ans int) {
	for num > 0 {
		if num&1 == 1 {
			ans++
		}
		num >>= 1
	}
	return
}
