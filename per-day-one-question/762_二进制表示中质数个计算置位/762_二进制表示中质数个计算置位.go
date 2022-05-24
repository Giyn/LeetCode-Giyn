/*
-------------------------------------
# @Time    : 2022/4/5 0:30:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 762_二进制表示中质数个计算置位.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	left := 10
	right := 15
	fmt.Println(countPrimeSetBits(left, right))
}

func countPrimeSetBits(left int, right int) (ans int) {
	isPrime := func(x int) bool {
		if x < 2 {
			return false
		}
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for i := left; i <= right; i++ {
		cnt := 0
		tmp := i
		for tmp > 0 {
			if tmp&1 == 1 {
				cnt++
			}
			tmp >>= 1
		}
		if isPrime(cnt) {
			ans++
		}
	}
	return
}
