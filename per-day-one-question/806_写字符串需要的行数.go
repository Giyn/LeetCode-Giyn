/*
-------------------------------------
# @Time    : 2022/4/12 14:31:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 806_写字符串需要的行数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(numberOfLines(widths, s))
}

func numberOfLines(widths []int, s string) []int {
	row, width := 1, 0
	for _, ch := range s {
		if width+widths[ch-'a'] > 100 {
			width = 0
			row++
		}
		width += widths[ch-'a']
	}
	return []int{row, width}
}
