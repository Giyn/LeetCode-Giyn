/*
-------------------------------------
# @Time    : 2022/4/19 23:27:22
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 821_字符的最短距离.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "loveleetcode"
	c := byte('e')
	fmt.Println(shortestToChar(s, c))
}

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)
	for i, last := 0, -n; i < n; i++ {
		ans[i] = n
		if s[i] == c {
			ans[i] = 0
			last = i
		} else {
			ans[i] = i - last
		}
	}
	for i, last := n-1, 2*n; i >= 0; i-- {
		if s[i] == c {
			last = i
		} else if v := last - i; v < ans[i] {
			ans[i] = v
		}
	}
	return ans
}
