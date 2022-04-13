/*
-------------------------------------
# @Time    : 2022/4/11 0:12:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 357_统计各位数字都不同的数字个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 4
	fmt.Println(countNumbersWithUniqueDigits(n))
}

func countNumbersWithUniqueDigits(n int) int {
	ans := 1
	for tmp, generate := 9, 9; n > 0; generate, n = generate-1, n-1 {
		ans += tmp
		tmp *= generate
	}
	return ans
}
