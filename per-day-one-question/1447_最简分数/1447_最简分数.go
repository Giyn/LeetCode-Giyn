/*
-------------------------------------
# @Time    : 2022/2/10 23:38:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1447_最简分数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 4
	fmt.Println(simplifiedFractions(n))
}

func simplifiedFractions(n int) (ans []string) {
	// hasCommonFactor := func(x, y int) bool {
	// 	for i := 2; i <= y; i++ {
	// 		if x%i == 0 && y%i == 0 {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }
	// for i := 2; i <= n; i++ {
	// 	for j := 1; j < i; j++ {
	// 		if !hasCommonFactor(i, j) {
	// 			ans = append(ans, fmt.Sprintf("%d/%d", j, i))
	// 		}
	// 	}
	// }
	// return

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	for denominator := 2; denominator <= n; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if gcd(denominator, numerator) == 1 {
				ans = append(ans, fmt.Sprintf("%d/%d", numerator, denominator))
			}
		}
	}
	return
}
