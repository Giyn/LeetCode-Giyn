/*
-------------------------------------
# @Time    : 2022/3/28 0:11:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 693_交替位二进制数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 5
	fmt.Println(hasAlternatingBits(n))
}

func hasAlternatingBits(n int) bool {
	pre := -1
	for n > 0 {
		if n&1 == pre {
			return false
		}
		pre = n & 1
		n >>= 1
	}
	return true
}
