/*
-------------------------------------
# @Time    : 2022/5/4 18:33:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1823_找出游戏的获胜者.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 5
	k := 2
	fmt.Println(findTheWinner(n, k))
}

func findTheWinner(n int, k int) (ans int) {
	for i := 2; i <= n; i++ {
		ans = (ans + k) % i
	}
	ans++
	return
}
