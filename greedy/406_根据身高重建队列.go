/*
-------------------------------------
# @Time    : 2021/12/19 10:16:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 406_根据身高重建队列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	people := [][]int{{7, 0}, {7, 1}, {6, 1}, {5, 0}, {5, 2}, {4, 4}}
	fmt.Println(reconstructQueue(people))
}

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	for _, info := range people {
		ans = append(ans, info)
		copy(ans[info[1]+1:], ans[info[1]:])
		ans[info[1]] = info
	}
	return
}
