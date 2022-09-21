/*
-------------------------------------
# @Time    : 2022/9/21 23:34:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_046_把数字翻译成字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 12258
	fmt.Println(translateNum(num))
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	if num >= 10 && num <= 25 {
		return 2
	}
	str := strconv.Itoa(num)
	dp := make([]int, len(str))
	dp[0] = 1
	if str[:2] >= "10" && str[:2] <= "25" { // 边界
		dp[1] = 2
	} else {
		dp[1] = 1
	}
	for i := 2; i < len(str); i++ {
		newNum := str[i-1 : i+1]
		if newNum >= "10" && newNum <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(str)-1]
}
