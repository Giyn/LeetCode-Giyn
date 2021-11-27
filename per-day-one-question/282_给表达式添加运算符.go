/*
-------------------------------------
# @Time    : 2021/10/16 8:55:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 282_给表达式添加运算符.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := "123"
	target := 6
	fmt.Println(addOperators(num, target))
}

func addOperators(num string, target int) (ans []string) {
	n := len(num)
	var backtrack func(expr []byte, i, res, mul int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)
		}
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				backtrack(expr, j+1, val, val)
			} else {
				expr[signIndex] = '+'
				backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, res-val, -val)
				expr[signIndex] = '*'
				backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return
}
