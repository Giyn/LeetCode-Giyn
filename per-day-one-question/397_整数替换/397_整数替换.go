/*
-------------------------------------
# @Time    : 2021/11/20 10:07:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 397_整数替换.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 8
	fmt.Println(integerReplacement(n))
}

func integerReplacement(n int) (ans int) {
	// DFS + 记忆化搜索
	memo := make(map[int]int)
	var _integerReplacement func(n int) int
	_integerReplacement = func(n int) (res int) {
		if n == 1 {
			return 0
		}
		if v, ok := memo[n]; ok {
			return v
		}
		defer func() {
			memo[n] = res
		}()
		if n%2 == 0 {
			return 1 + _integerReplacement(n/2)
		}
		return 2 + min(_integerReplacement(n/2), _integerReplacement(n/2+1))
	}
	return _integerReplacement(n)

	// DFS
	//if n == 1 {
	//	return 0
	//}
	//if n%2 == 0 {
	//	return 1 + integerReplacement(n/2)
	//}
	//return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))

	//for n != 1 {
	//	switch {
	//	case n%2 == 0:
	//		ans++
	//		n /= 2
	//	case n%4 == 1:
	//		ans += 2
	//		n /= 2
	//	case n == 3:
	//		ans += 2
	//		n = 1
	//	default:
	//		ans += 2
	//		n = n/2 + 1
	//	}
	//}
	//return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
