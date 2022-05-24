/*
-------------------------------------
# @Time    : 2022/3/22 9:26:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2038_如果相邻两个颜色均相同则删除当前颜色.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	colors := "AAABABB"
	fmt.Println(winnerOfGame(colors))
}

func winnerOfGame(colors string) bool {
	var count int
	for i := 0; i < len(colors)-2; i++ {
		if colors[i:i+3] == "AAA" {
			count++
		} else if colors[i:i+3] == "BBB" {
			count--
		}
	}
	return count > 0
}
