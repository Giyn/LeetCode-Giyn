/*
-------------------------------------
# @Time    : 2022/3/25 12:11:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 172_阶乘后的零.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 3
	fmt.Println(trailingZeroes(n))
}

func trailingZeroes(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}
