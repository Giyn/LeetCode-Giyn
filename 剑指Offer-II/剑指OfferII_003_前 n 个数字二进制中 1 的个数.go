/*
-------------------------------------
# @Time    : 2022/3/3 9:45:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_003_前 n 个数字二进制中 1 的个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	n := 5
	fmt.Println(countBits(n))
}

func countBits(n int) (ans []int) {
	// dp[i] = dp[i-1]+1，当i为奇数
	// dp[i] = dp[i/2]，当i为偶数
	// 合并为 dp[i] = dp[i/2] + i % 2
	// i / 2 可以通过 i >> 1 得到；
	// i % 2 可以通过 i & 1 得到；
	ans = make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = ans[i>>1] + (i & 1)
	}
	return

	//for i := 0; i <= n; i++ {
	//	b := fmt.Sprintf("%b", i)
	//	ans = append(ans, strings.Count(b, "1"))
	//}
	//return
}
