/*
-------------------------------------
# @Time    : 2021/10/24 20:28:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 638_大礼包.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	price := []int{2, 5}
	special := [][]int{{3, 0, 5}, {1, 2, 10}}
	needs := []int{3, 2}
	fmt.Println(shoppingOffers(price, special, needs))
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)

	var filterSpecial [][]int
	for _, s := range special {
		totalCount, totalPrice := 0, 0
		for i, c := range s[:n] {
			totalCount += c
			totalPrice += c * price[i]
		}
		if totalPrice > 0 && totalPrice > s[n] {
			filterSpecial = append(filterSpecial, s)
		}
	}

	dp := map[string]int{}
	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minPrice int) {
		fmt.Println(curNeeds)
		if res, has := dp[string(curNeeds)]; has {
			return res
		}
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p
		}
		nextNeeds := make([]byte, n)
	outer:
		for _, s := range filterSpecial {
			for i, need := range curNeeds {
				if need < byte(s[i]) {
					continue outer
				}
				nextNeeds[i] = need - byte(s[i])
			}
			minPrice = min638(minPrice, dfs(nextNeeds)+s[n])
		}
		dp[string(curNeeds)] = minPrice
		return
	}

	curNeeds := make([]byte, n)
	for i, need := range needs {
		curNeeds[i] = byte(need)
	}
	return dfs(curNeeds)
}

func min638(x, y int) int {
	if x < y {
		return x
	}
	return y
}
