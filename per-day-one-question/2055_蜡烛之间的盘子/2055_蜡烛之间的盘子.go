/*
-------------------------------------
# @Time    : 2022/3/8 8:45:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2055_蜡烛之间的盘子.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "**|**|***|"
	queries := [][]int{{2, 5}, {5, 9}}
	fmt.Println(platesBetweenCandles(s, queries))
}

func platesBetweenCandles(s string, queries [][]int) (ans []int) {
	n := len(s)
	pre, left, right := make([]int, n+1), make([]int, n), make([]int, n)
	for i, j, l, r := 0, n-1, -1, -1; i < n; i, j = i+1, j-1 {
		if s[i] == '*' {
			pre[i+1] = pre[i] + 1
		} else {
			pre[i+1] = pre[i]
			l = i
		}
		if s[j] == '|' {
			r = j
		}
		left[i] = l
		right[j] = r
	}
	ans = make([]int, len(queries))
	for i, query := range queries {
		if left[query[1]] >= 0 && right[query[0]] >= 0 && left[query[1]] > right[query[0]] {
			ans[i] = pre[left[query[1]]] - pre[right[query[0]]]
		}
	}
	return
}
