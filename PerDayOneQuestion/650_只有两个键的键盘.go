/*
-------------------------------------
# @Time    : 2021/10/12 17:02:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 650_只有两个键的键盘.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main () {
	n := 3
	fmt.Println(minSteps(n))
}

func minSteps(n int) int {
	ans := 0
	for i := 2; i * i <= n; i++ {
		for n % i == 0 {
			ans += i
			n /= i
		}
	}
	if n != 1 {
		ans += n
	}
	return ans
}
