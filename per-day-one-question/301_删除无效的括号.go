/*
-------------------------------------
# @Time    : 2021/10/27 10:53:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 301_删除无效的括号.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "(a)())()"
	//s := ")("
	fmt.Println(removeInvalidParentheses(s))
}

func removeInvalidParentheses(str string) (ans []string) {
	isValid := func(str string) bool {
		cnt := 0
		for _, ch := range str {
			if ch == '(' {
				cnt++
			} else if ch == ')' {
				cnt--
				if cnt < 0 {
					return false
				}
			}
		}
		return cnt == 0
	}
	var dfs func(str string, start, l, r int, ans *[]string)
	dfs = func(str string, start, l, r int, ans *[]string) {
		if l == 0 && r == 0 {
			if isValid(str) {
				*ans = append(*ans, str)
			}
			return
		}
		for i := start; i < len(str); i++ {
			// 多种相同情况只需去除一次
			if i != start && str[i] == str[i-1] {
				continue
			}
			if str[i] == ')' && r > 0 {
				cur := str
				cur = cur[:i] + cur[i+1:]
				dfs(cur, i, l, r-1, ans)
			} else if str[i] == '(' && l > 0 {
				cur := str
				cur = cur[:i] + cur[i+1:]
				dfs(cur, i, l-1, r, ans)
			}
		}
	}
	l, r := 0, 0
	for _, v := range str {
		if v == '(' {
			l++
		} else if v == ')' {
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}
	dfs(str, 0, l, r, &ans)
	return
}
