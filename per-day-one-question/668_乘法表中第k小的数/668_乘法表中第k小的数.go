/*
-------------------------------------
# @Time    : 2022/5/18 0:58:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 668_乘法表中第k小的数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := 3
	n := 3
	k := 5
	fmt.Println(findKthNumber(m, n, k))
}

func findKthNumber(m, n, k int) int {
	return sort.Search(m*n, func(x int) bool {
		count := x / n * n
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		return count >= k
	})
}
