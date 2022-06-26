/*
-------------------------------------
# @Time    : 2022/6/27 3:21:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 043_字符串相乘.go
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
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			sum := tmp[i+j+1] + n1*n2
			tmp[i+j+1] = sum % 10
			tmp[i+j] += sum / 10
		}
	}
	ans := strings.Builder{}
	for i := 0; i < len(tmp); i++ {
		if i == 0 && tmp[i] == 0 {
			continue
		}
		ans.WriteString(strconv.Itoa(tmp[i]))
	}
	return ans.String()
}
