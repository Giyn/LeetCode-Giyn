/*
-------------------------------------
# @Time    : 2022/2/7 19:51:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1405_最长快乐字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	a, b, c := 1, 1, 7
	fmt.Println(longestDiverseString(a, b, c))
}

func longestDiverseString(a, b, c int) string {
	var ans []byte
	cnt := []struct {
		num int
		ch  byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	for {
		sort.Slice(cnt, func(i, j int) bool { return cnt[i].num > cnt[j].num })
		hasNext := false
		for i, pair := range cnt {
			if pair.num <= 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-2] == pair.ch && ans[m-1] == pair.ch {
				continue
			}
			hasNext = true
			ans = append(ans, pair.ch)
			cnt[i].num--
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}
