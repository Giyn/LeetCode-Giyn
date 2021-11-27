/*
-------------------------------------
# @Time    : 2021/11/4 0:23:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 367_有效的完全平方数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := 104976
	fmt.Println(isPerfectSquare(num))
}

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 1, num
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
