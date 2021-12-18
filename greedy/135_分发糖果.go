/*
-------------------------------------
# @Time    : 2021/12/18 10:09:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 135_分发糖果.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	ratings := []int{1, 2, 2}
	fmt.Println(candy(ratings))
}

func candy(ratings []int) (ans int) {
	n := len(ratings)
	candyAllocated := make([]int, n)
	for i := range candyAllocated {
		candyAllocated[i] = 1
	}
	for i := 0; i < n-1; i++ {
		if ratings[i+1] > ratings[i] {
			candyAllocated[i+1] = candyAllocated[i] + 1
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			if candyAllocated[i]+1 > candyAllocated[i-1] {
				candyAllocated[i-1] = candyAllocated[i] + 1
			}
		}
	}
	for i := 0; i < n; i++ {
		ans += candyAllocated[i]
	}
	return
}
