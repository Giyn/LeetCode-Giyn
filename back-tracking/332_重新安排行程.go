/*
-------------------------------------
# @Time    : 2021/12/8 22:11:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 332_重新安排行程.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{{"EZE", "TIA"}, {"EZE", "AXA"}, {"AUA", "EZE"}, {"EZE", "JFK"}, {"JFK", "ANU"}, {"JFK", "ANU"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"TIA", "JFK"}, {"ANU", "EZE"}, {"ANU", "EZE"}, {"TIA", "AUA"}}
	fmt.Println(findItinerary(tickets))
}

func findItinerary(tickets [][]string) []string {
	var ans []string
	var mpSort = make(map[string][]string)
	var mp = make(map[string]map[string]int)
	var ticketNum = len(tickets)

	var backtrack func() bool
	backtrack = func() bool {
		if len(ans) == ticketNum+1 {
			return true
		}
		last := ans[len(ans)-1]
		for k, v := range mp[last] {
			if v > 0 {
				ans = append(ans, k)
				mp[last][k]--
				if backtrack() {
					return true
				}
				ans = ans[:len(ans)-1]
				mp[last][k]++
			}
		}
		return false
	}
	for k := range mp {
		delete(mp, k)
	}
	for k := range mpSort {
		delete(mpSort, k)
	}
	for _, ticket := range tickets {
		if _, ok := mp[ticket[0]]; !ok {
			mp[ticket[0]] = make(map[string]int)
		}
		mpSort[ticket[0]] = append(mpSort[ticket[0]], ticket[1])
	}
	for k := range mpSort {
		sort.Strings(mpSort[k])
	}
	for k, v := range mpSort {
		for i := 0; i < len(v); i++ {
			mp[k][v[i]]++
		}
	}
	ans = append(ans, "JFK")
	backtrack()
	return ans
}
