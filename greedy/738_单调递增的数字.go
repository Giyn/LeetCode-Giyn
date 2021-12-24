/*
-------------------------------------
# @Time    : 2021/12/24 15:49:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 738_单调递增的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 332
	fmt.Println(monotoneIncreasingDigits(n))
}

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	b := []byte(s)
	length := len(b)
	if length <= 1 {
		return n
	}
	for i := length - 1; i > 0; i-- {
		if b[i-1] > b[i] {
			b[i-1] -= 1
			for j := i; j < length; j++ {
				b[j] = '9'
			}
		}
	}
	ans, _ := strconv.Atoi(string(b))
	return ans
}
