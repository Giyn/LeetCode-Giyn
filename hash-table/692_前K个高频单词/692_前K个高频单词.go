/*
-------------------------------------
# @Time    : 2022/4/7 17:00:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 692_前K个高频单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))
}

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int, len(words))
	for _, word := range words {
		mp[word]++
	}
	var ans []string
	for key := range mp {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(i, j int) bool {
		return mp[ans[i]] > mp[ans[j]] || mp[ans[i]] == mp[ans[j]] && ans[i] < ans[j]
	})
	return ans[:k]
}
