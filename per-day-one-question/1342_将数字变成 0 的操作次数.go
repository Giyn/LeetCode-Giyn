/*
-------------------------------------
# @Time    : 2022/1/31 12:26:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1342_将数字变成 0 的操作次数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := 14
	fmt.Println(numberOfSteps(num))
}

func numberOfSteps(num int) (ans int) {
	for num > 0 {
		if num%2 == 1 {
			num--
		} else {
			num >>= 1
		}
		ans++
	}
	return
}
