/*
-------------------------------------
# @Time    : 2022/1/5 1:48:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 507_完美数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := 4
	fmt.Println(checkPerfectNumber(num))
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sum += d
			if d*d < num {
				sum += num / d
			}
		}
	}
	return sum == num
}
