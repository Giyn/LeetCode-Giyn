/*
-------------------------------------
# @Time    : 2022/1/19 13:32:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 509_斐波那契数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 3
	fmt.Println(fib(n))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		tmp := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = tmp
	}
	return dp[1]
}
