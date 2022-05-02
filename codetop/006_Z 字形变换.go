/*
-------------------------------------
# @Time    : 2022/4/29 16:29:03
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
	if numRows == 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	row := 0
	cycle := 2*numRows - 2
	for i := 0; i < len(s); i++ {
		matrix[row] = append(matrix[row], s[i])
		if i%cycle < numRows-1 {
			row++
		} else {
			row--
		}
	}
	return string(bytes.Join(matrix, nil))
}
