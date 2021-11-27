/*
-------------------------------------
# @Time    : 2021/11/8 0:08:08
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 299_猜数字游戏.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	secret := "1807"
	guess := "7810"
	fmt.Println(getHint(secret, guess))
}

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	recordS, recordG := [10]int{}, [10]int{}
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			recordS[secret[i]-'0']++
			recordG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min299(recordS[i], recordG[i])
	}
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(bulls))
	builder.WriteString("A")
	builder.WriteString(strconv.Itoa(cows))
	builder.WriteString("B")

	return builder.String()
}

func min299(x, y int) int {
	if x < y {
		return x
	}
	return y
}
