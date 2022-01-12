/*
-------------------------------------
# @Time    : 2022/1/10 0:26:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 306_累加数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	num := "112358"
	fmt.Println(isAdditiveNumber(num))
}

func isAdditiveNumber(num string) bool {
	var add func(a, b string) string
	add = func(a, b string) string {
		var res []byte
		one := 0
		for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; {
			curA, curB := 0, 0
			if i >= 0 {
				curA = int(a[i] - '0')
				i--
			}
			if j >= 0 {
				curB = int(b[j] - '0')
				j--
			}
			cur := curA + curB + one
			one = cur / 10
			res = append(res, byte(cur%10)+'0')
		}
		if one == 1 {
			res = append(res, '1')
		}
		for i, n := 0, len(res); i < n/2; i++ {
			res[i], res[n-1-i] = res[n-1-i], res[i]
		}
		return string(res)
	}

	var dfs func(num string, first, second int) bool
	dfs = func(num string, first, second int) bool {
		n, last := len(num), 0
		for second < n {
			if (num[last] == '0' && first > last+1) || (num[first] == '0' && second > first+1) {
				return false
			}
			s := add(num[last:first], num[first:second])
			if second+len(s) > n || num[second:second+len(s)] != s {
				return false
			}
			last, first, second = first, second, second+len(s)
		}
		return true
	}

	for i := 1; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if dfs(num, i, j) {
				return true
			}
		}
	}
	return false
}
