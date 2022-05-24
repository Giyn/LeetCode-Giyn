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
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(findItinerary(tickets))
}

func findItinerary(tickets [][]string) (ans []string) {
	mp := map[string][]string{}

	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		mp[src] = append(mp[src], dst)
	}
	for key := range mp {
		sort.Strings(mp[key])
	}
	var dfs func(cur string)
	dfs = func(cur string) {
		for {
			if v, ok := mp[cur]; !ok || len(v) == 0 {
				break
			}
			tmp := mp[cur][0]
			mp[cur] = mp[cur][1:]
			dfs(tmp)
		}
		ans = append(ans, cur)
	}
	dfs("JFK")
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return
}
