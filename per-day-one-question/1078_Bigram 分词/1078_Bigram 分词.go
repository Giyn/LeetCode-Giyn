/*
-------------------------------------
# @Time    : 2021/12/26 2:01:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1078_Bigram 分词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "alice is a good girl she is a good student"
	first := "a"
	second := "good"
	fmt.Println(findOcurrences(text, first, second))
}

func findOcurrences(text string, first string, second string) (ans []string) {
	textSlice := strings.Split(text, " ")
	for i := 2; i < len(textSlice); i++ {
		if textSlice[i-1] == second && textSlice[i-2] == first {
			ans = append(ans, textSlice[i])
		}
	}
	return
}
