/*
-------------------------------------
# @Time    : 2021/10/13 9:47:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 412_Fizz Buzz.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = strconv.Itoa(i + 1)
	}
	for i, v := range ans {
		val, _ := strconv.Atoi(v)
		if val%3 == 0 && val%5 != 0 {
			ans[i] = "Fizz"
		}
		if val%3 != 0 && val%5 == 0 {
			ans[i] = "Buzz"
		}
		if val%3 == 0 && val%5 == 0 {
			ans[i] = "FizzBuzz"
		}
	}
	return ans
}

//func fizzBuzz(n int) (ans []string) {
//	for i := 1; i <= n; i++ {
//		sb := strings.Builder{}
//		if i%3 == 0 {
//			sb.WriteString("Fizz")
//		}
//		if i%5 == 0 {
//			sb.WriteString("Buzz")
//		}
//		if sb.Len() == 0 {
//			sb.WriteString(strconv.Itoa(i))
//		}
//		ans = append(ans, sb.String())
//	}
//	return
//}
