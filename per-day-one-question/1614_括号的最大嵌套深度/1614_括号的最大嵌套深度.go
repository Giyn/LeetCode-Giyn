/*
-------------------------------------
# @Time    : 2022/1/7 1:36:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1614_括号的最大嵌套深度.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))
}

func maxDepth(s string) (ans int) {
	tmp := 0
	for _, ch := range s {
		if ch == '(' {
			tmp++
		} else if ch == ')' {
			tmp--
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return
}
