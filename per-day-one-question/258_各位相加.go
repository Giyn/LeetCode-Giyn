/*
-------------------------------------
# @Time    : 2022/3/3 9:12:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 258_各位相加.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := 38
	fmt.Println(addDigits(num))
}

func addDigits(num int) (ans int) {
	if num < 10 {
		return num
	}
	for num > 0 {
		ans += num % 10
		num /= 10
	}
	if ans >= 10 {
		ans = addDigits(ans)
	}
	return
}
