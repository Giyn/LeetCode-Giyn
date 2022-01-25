/*
-------------------------------------
# @Time    : 2022/1/25 0:19:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1688_比赛中的配对次数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 14
	fmt.Println(numberOfMatches(n))
}

func numberOfMatches(n int) (ans int) {
	//for n >= 2 {
	//	ans += n / 2
	//	if n%2 == 1 {
	//		n = n/2 + 1
	//	} else {
	//		n /= 2
	//	}
	//}
	//return
	return n - 1
}
