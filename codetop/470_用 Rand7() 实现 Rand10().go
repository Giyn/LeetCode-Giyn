/*
-------------------------------------
# @Time    : 2022/4/16 15:03:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 470_用 Rand7() 实现 Rand10().go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(rand10())
	}
}

func rand10() (ans int) {
	for {
		ans = (rand7()-1)*7 + rand7()
		if ans <= 40 {
			return 1 + ans%10
		}
		// 41-49
		ans = (ans-40-1)*7 + rand7()
		if ans <= 60 {
			return 1 + ans%10
		}
		// 61-63
		ans = (ans-60-1)*7 + rand7()
		if ans <= 20 {
			return 1 + ans%10
		}
	}
}

func rand7() int {
	return 1 + rand.Intn(7)
}
