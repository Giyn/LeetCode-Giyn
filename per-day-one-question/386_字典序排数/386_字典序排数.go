/*
-------------------------------------
# @Time    : 2022/4/18 0:14:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 386_字典序排数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 100
	fmt.Println(lexicalOrder(n))
}

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			if num == n {
				break
			}
			for num%10 == 9 {
				num /= 10
			}
			num++
		}
	}
	return ans
}
