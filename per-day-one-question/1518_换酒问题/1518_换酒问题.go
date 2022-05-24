/*
-------------------------------------
# @Time    : 2021/12/17 0:31:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1518_换酒问题.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numBottles := 2
	numExchange := 3
	fmt.Println(numWaterBottles(numBottles, numExchange))
}

func numWaterBottles(numBottles int, numExchange int) (ans int) {
	ans = numBottles
	var remainBottles = numBottles
	for remainBottles/numExchange != 0 {
		newBottles := remainBottles / numExchange
		remainBottles %= numExchange
		ans += newBottles
		remainBottles += newBottles
	}
	return
}
