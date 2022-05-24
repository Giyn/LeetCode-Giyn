/*
-------------------------------------
# @Time    : 2022/5/23 0:03:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 405_数字转换为十六进制数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	num := -1
	fmt.Println(toHex(num))
}

func toHex(num int) (ans string) {
	if num == 0 {
		return "0"
	}
	const CONV = "0123456789abcdef"
	for num != 0 && len(ans) < 8 {
		tmp := num & 15 // 低4位
		ans = string(CONV[tmp]) + ans
		num >>= 4
	}
	return
}
