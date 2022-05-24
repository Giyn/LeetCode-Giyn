/*
-------------------------------------
# @Time    : 2022/3/12 17:24:54
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_033_变位词组.go
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
	for _, str := range strs {
		record := [26]int{}
		for i := 0; i < len(str); i++ {
			record[str[i]-'a']++
		}
		mp[record] = append(mp[record], str)
	}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return
}
