/*
-------------------------------------
# @Time    : 2022/3/1 19:26:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 022_括号生成.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 4
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) (ans []string) {
	var getParenthesis func(string, int, int)
	getParenthesis = func(str string, lRemain, rRemain int) {
		if lRemain == 0 && rRemain == 0 {
			ans = append(ans, str)
			return
		}
		if lRemain == rRemain {
			getParenthesis(str+"(", lRemain-1, rRemain)
		} else if lRemain < rRemain {
			if lRemain > 0 {
				getParenthesis(str+"(", lRemain-1, rRemain)
			}
			getParenthesis(str+")", lRemain, rRemain-1)
		}
	}
	getParenthesis("", n, n)
	return
}
