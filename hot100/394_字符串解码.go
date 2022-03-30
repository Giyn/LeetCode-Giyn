/*
-------------------------------------
# @Time    : 2022/3/30 10:48:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 394_字符串解码.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			// 处理字母
			var letter []byte
			for len(stack) > 0 && unicode.IsLetter(rune(stack[len(stack)-1])) {
				letter = append([]byte{stack[len(stack)-1]}, letter...)
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // 去除‘[’
			// 处理数字
			var digit []byte
			for len(stack) > 0 && unicode.IsDigit(rune(stack[len(stack)-1])) {
				digit = append([]byte{stack[len(stack)-1]}, digit...)
				stack = stack[:len(stack)-1]
			}
			for count, _ := strconv.Atoi(string(digit)); count > 0; count-- {
				for _, ch := range letter {
					stack = append(stack, ch)
				}
			}
		}
	}
	return string(stack)
}
