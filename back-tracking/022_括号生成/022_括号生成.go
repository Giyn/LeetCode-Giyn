/*
-------------------------------------
# @Time    : 2022/3/19 23:27:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 022_括号生成.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) (ans []string) {
	var dfs func(str string, lRemain, rRemain int)
	dfs = func(str string, lRemain, rRemain int) {
		if len(str) == 2*n {
			ans = append(ans, str)
			return
		}
		if lRemain > 0 {
			dfs(str+"(", lRemain-1, rRemain)
		}
		if lRemain < rRemain {
			dfs(str+")", lRemain, rRemain-1)
		}
	}
	dfs("", n, n)
	return
}
