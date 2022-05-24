/*
-------------------------------------
# @Time    : 2022/3/1 9:51:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 006_Z 字形变换.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	mat := make([][]byte, numRows)
	t, x := numRows*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < numRows-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}
