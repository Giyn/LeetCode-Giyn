/*
-------------------------------------
# @Time    : 2022/2/4 17:15:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1414_和为 K 的最少斐波那契数字数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	k := 7
	fmt.Println(findMinFibonacciNumbers(k))
}

func findMinFibonacciNumbers(k int) (ans int) {
	f := []int{1, 1}
	for f[len(f)-1] < k {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}
	for i := len(f) - 1; k > 0; i-- {
		if f[i] <= k {
			k -= f[i]
			ans++
		}
	}
	return
}
