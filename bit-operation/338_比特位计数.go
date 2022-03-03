/*
-------------------------------------
# @Time    : 2022/3/3 10:13:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 338_比特位计数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 5
	fmt.Println(countBits(n))
}

func countBits(n int) (ans []int) {
	ans = make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return

	// for i := 0; i <= n; i++ {
	// 	b := fmt.Sprintf("%b", i)
	// 	ans = append(ans, strings.Count(b, "1"))
	// }
	// return
}
