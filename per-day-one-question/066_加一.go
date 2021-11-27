/*
-------------------------------------
# @Time    : 2021/10/22 0:40:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 066_加一.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	// digits := []int{1, 2, 3}
	digits := []int{9, 9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	n := len(digits)
	flag := false

	for i := range digits {
		if digits[n-i-1] == 9 {
			flag = true
			digits[n-i-1] = 0
		} else {
			flag = false
			digits[n-i-1]++
			return digits
		}
	}
	if flag {
		digits = append(digits, 1)
		copy(digits[1:n+1], digits[:n])
		digits[0] = 1
		// or
		// digits = append([]int{1}, digits...)
	}
	return digits
}
