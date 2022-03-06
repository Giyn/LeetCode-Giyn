/*
-------------------------------------
# @Time    : 2022/3/6 9:47:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2100_适合打劫银行的日子.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	security := []int{1, 2, 5, 4, 1, 0, 2, 4, 5, 3, 1, 2, 4, 3, 2, 4, 8}
	time := 2
	fmt.Println(goodDaysToRobBank(security, time))
}

func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1] + 1
		}
		if security[n-i] >= security[n-1-i] {
			right[n-1-i] = right[n-i] + 1
		}
	}
	for i := 0; i < n; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}
