/*
-------------------------------------
# @Time    : 2022/5/19 0:52:51
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_013_机器人的运动范围.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	m := 2
	n := 3
	k := 1
	fmt.Println(movingCount(m, n, k))
}

func movingCount(m int, n int, k int) (ans int) {
	digitSum := func(x int) (ans int) {
		for x > 0 {
			ans += x % 10
			x /= 10
		}
		return
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	ans = 1 // 起点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 || digitSum(i)+digitSum(j) > k {
				continue
			}
			if i >= 1 {
				visited[i][j] = visited[i-1][j] || visited[i][j]
			}
			if j >= 1 {
				visited[i][j] = visited[i][j-1] || visited[i][j]
			}
			if visited[i][j] {
				ans++
			}
		}
	}
	return
}
