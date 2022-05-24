/*
-------------------------------------
# @Time    : 2022/2/26 9:56:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 537_复数乘法.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num1 := "1+-1i"
	num2 := "1+-1i"
	fmt.Println(complexNumberMultiply(num1, num2))
}

func complexNumberMultiply(num1 string, num2 string) string {
	i1 := strings.IndexByte(num1, '+')
	i2 := strings.IndexByte(num2, '+')
	re1, _ := strconv.Atoi(num1[:i1])
	re2, _ := strconv.Atoi(num2[:i2])
	im1, _ := strconv.Atoi(num1[i1+1 : len(num1)-1])
	im2, _ := strconv.Atoi(num2[i2+1 : len(num2)-1])

	return fmt.Sprintf("%d+%di", re1*re2-im1*im2, re1*im2+re2*im1)
}
