/*
-------------------------------------
# @Time    : 2021/11/1 13:31:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 049_字母异位词分组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) (ans [][]string) {
	mp := make(map[[26]int][]string)
	for _, word := range strs {
		record := [26]int{}
		for i := 0; i < len(word); i++ {
			record[word[i]-'a']++
		}
		mp[record] = append(mp[record], word)
	}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return
}
