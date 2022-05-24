/*
-------------------------------------
# @Time    : 2022/4/25 16:28:04
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
	var dfs func(str string, left, right int)
	dfs = func(str string, left, right int) {
		if len(str) == 2*n {
			ans = append(ans, str)
			return
		}
		if left > 0 {
			dfs(str+"(", left-1, right)
		}
		if left < right {
			dfs(str+")", left, right-1)
		}
	}
	dfs("", n, n)
	return
}
