/*
-------------------------------------
# @Time    : 2021/12/5 9:13:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 372_超级次方.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	a := 2147483647
	b := []int{2, 0, 0}
	fmt.Println(superPow(a, b))
}

func superPow(a int, b []int) int {
	var MOD = 1337
	var quickPow func(a, b int) int
	var dfs func(a int, b []int) int
	quickPow = func(a, b int) int {
		ans := 1
		a %= MOD
		for b != 0 {
			if b&1 == 1 {
				ans = ans * a % MOD
			}
			a = a * a % MOD
			b >>= 1
		}
		return ans
	}
	dfs = func(a int, b []int) int {
		if len(b) == 0 || a == 1 {
			return 1
		}
		// (a * b) % p = (a % p * b % p) % p
		return quickPow(dfs(a, b[:len(b)-1]), 10) * quickPow(a, b[len(b)-1]) % MOD
	}
	a %= MOD
	return dfs(a, b)
}
