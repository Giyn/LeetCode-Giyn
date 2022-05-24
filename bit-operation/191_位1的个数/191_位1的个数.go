/*
-------------------------------------
# @Time    : 2022/3/31 23:13:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 191_位1的个数.go
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
