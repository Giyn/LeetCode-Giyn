/*
-------------------------------------
# @Time    : 2021/11/28 1:21:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 077_组合.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

var ans [][]int
var track []int

func combine(n int, k int) [][]int {
	ans = [][]int{}
	track = []int{}
	if n <= 0 || k <= 0 || k > n {
		return ans
	}
	backtrack(n, k, 1)
	return ans
}

func backtrack(n, k, start int) {
	if len(track) == k {
		ans = append(ans, append([]int{}, track...))
		return
	}
	// 剪枝
	if len(track)+n-start+1 < k {
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1)
		track = track[:len(track)-1]
	}
}
