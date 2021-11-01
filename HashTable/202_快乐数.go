/*
-------------------------------------
# @Time    : 2021/11/1 22:39:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 202_快乐数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	fmt.Println(isHappy(7))
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	mp := make(map[int]bool)
	for {
		sum := 0
		for n > 0 {
			bit := n % 10
			sum += bit * bit
			n /= 10
		}
		if sum == 1 {
			return true
		}
		n = sum
		if ok, _ := mp[sum]; ok {
			return false
		}
		mp[sum] = true
	}
}
