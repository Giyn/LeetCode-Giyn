/*
-------------------------------------
# @Time    : 2022/3/31 0:03:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 728_自除数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	left := 1
	right := 22
	fmt.Println(selfDividingNumbers(left, right))
}

func selfDividingNumbers(left int, right int) (ans []int) {
	isDividingNumber := func(num int) bool {
		for tmp := num; tmp > 0; tmp /= 10 {
			bit := tmp % 10
			if bit == 0 || num%(bit) != 0 {
				return false
			}
		}
		return true
	}
	for left <= right {
		if isDividingNumber(left) {
			ans = append(ans, left)
		}
		left++
	}
	return
}
