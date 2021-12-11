/*
-------------------------------------
# @Time    : 2021/12/12 3:08:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 709_转换成小写字母.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, World!"
	fmt.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	sb := strings.Builder{}
	for _, ch := range s {
		if ch >= 65 && ch <= 90 {
			ch |= 32
		}
		sb.WriteRune(ch)
	}
	return sb.String()
	//return strings.ToLower(s)
}
