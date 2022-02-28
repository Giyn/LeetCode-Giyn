/*
-------------------------------------
# @Time    : 2022/2/28 22:39:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1601_最多可达成的换楼请求数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 5
	requests := [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}
	fmt.Println(maximumRequests(n, requests))
}

func maximumRequests(n int, requests [][]int) (ans int) {
out:
	for i, m := 1, len(requests); i < 1<<m; i++ {
		cur, cnts := 0, make([]int, n)
		for j := 0; j < m; j++ {
			if (1<<j)&i > 0 {
				cnts[requests[j][0]]++
				cnts[requests[j][1]]--
				cur++
			}
		}
		for j := 0; j < n; j++ {
			if cnts[j] != 0 {
				continue out
			}
		}
		if cur > ans {
			ans = cur
		}
	}
	return
}
