/*
-------------------------------------
# @Time    : 2022/1/15 12:07:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1716_计算力扣银行的钱.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 1000
	fmt.Println(totalMoney(n))
}

func totalMoney(n int) (ans int) {
	totalWeek := n / 7
	remain := n % 7
	ans += totalWeek*28 + (totalWeek-1)*totalWeek/2*7
	for remain > 0 {
		ans += totalWeek + 1
		remain--
		totalWeek++
	}
	return
}
