/*
-------------------------------------
# @Time    : 2022/4/14 13:59:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1672_最富有客户的资产总量.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts))
}

func maximumWealth(accounts [][]int) (ans int) {
	for _, account := range accounts {
		tmp := 0
		for _, v := range account {
			tmp += v
		}
		if tmp > ans {
			ans = tmp
		}
	}
	return
}
