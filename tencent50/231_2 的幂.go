/*
-------------------------------------
# @Time    : 2022/3/30 23:55:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 231_2 的幂.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 9
	fmt.Println(isPowerOfTwo(n))
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for cnt := 0; n > 0; n >>= 1 {
		if n&1 == 1 {
			cnt++
		}
		if cnt == 2 {
			return false
		}
	}
	return true
}
