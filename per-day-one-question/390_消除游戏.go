/*
-------------------------------------
# @Time    : 2022/1/5 12:37:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 390_消除游戏.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 9
	fmt.Println(lastRemaining(n))
}

func lastRemaining(n int) int {
	a1 := 1
	k, cnt, step := 0, n, 1
	for cnt > 1 {
		if k%2 == 0 {
			a1 += step
		} else {
			if cnt%2 == 1 {
				a1 += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return a1
}
