/*
-------------------------------------
# @Time    : 2021/12/19 9:10:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 997_找到小镇的法官.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 1
	//trust := [][]int{{1, 3}, {2, 3}}
	var trust [][]int
	fmt.Println(findJudge(n, trust))
}

func findJudge(n int, trust [][]int) int {
	mp := make([]int, n+1)
	for _, v := range trust {
		mp[v[1]]++
		mp[v[0]]--
	}
	for i := 1; i <= n; i++ {
		if mp[i] == n-1 {
			return i
		}
	}
	return -1
}
