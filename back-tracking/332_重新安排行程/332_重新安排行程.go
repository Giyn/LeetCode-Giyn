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

func findItinerary(tickets [][]string) (ans []string) {
	var d = map[string][]string{}
	var dfs func(string)
	dfs = func(f string) {
		for len(d[f]) > 0 {
			v := d[f][0]
			d[f] = d[f][1:]
			dfs(v)
		}
		ans = append([]string{f}, ans...)
	}
	for _, v := range tickets {
		d[v[0]] = append(d[v[0]], v[1])
	}
	for _, v := range d {
		sort.Strings(v)
	}
	ans = []string{}
	dfs("JFK")
	return ans
}
