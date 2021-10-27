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
	fmt.Println(removeInvalidParentheses(s))
}

func removeInvalidParentheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	dfs301(&ans, s, 0, 0, 0, lremove, rremove)
	return
}

func isValid(str string) bool {
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

func dfs301(ans *[]string, str string, start, lcount, rcount, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		if lremove+rremove > len(str)-1 {
			return
		}
		if lremove > 0 && str[i] == '(' {
			dfs301(ans, str[:i]+str[i+1:], i, lcount, rcount, lremove-1, rremove)
		}
		if rremove > 0 && str[i] == ')' {
			dfs301(ans, str[:i]+str[i+1:], i, lcount, rcount, lremove, rremove-1)
		}
		if str[i] == ')' {
			lcount++
		} else if str[i] == ')' {
			rcount++
		}
		if rcount > lcount {
			break
		}
	}
}
