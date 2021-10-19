/*
-------------------------------------
# @Time    : 2021/10/19 8:58:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 292_Nim 游戏.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 4
	fmt.Println(canWinNim(n))
}

func canWinNim(n int) bool {
	return n%4 != 0
}
