/*
-------------------------------------
# @Time    : 2021/10/18 13:33:08
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 476_数字的补数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := 5
	fmt.Println(findComplement(num))
}

func findComplement(num int) int {
	i := 1
	for i <= num {
		i <<= 1
	}
	return num ^ (i - 1)
}
