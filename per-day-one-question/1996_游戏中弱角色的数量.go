/*
-------------------------------------
# @Time    : 2022/1/28 0:19:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1996_游戏中弱角色的数量.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	properties := [][]int{{7, 9}, {10, 7}, {6, 9}, {10, 4}, {7, 5}, {7, 10}}
	fmt.Println(numberOfWeakCharacters(properties))
}

func numberOfWeakCharacters(properties [][]int) (ans int) {
	// 排序
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})
	max := -1 // 记录最大防御值
	for i := 0; i < len(properties); i++ {
		if max > properties[i][1] {
			ans++
		} else if max < properties[i][1] {
			max = properties[i][1]
		}
	}
	return

	// 单调栈
	//sort.Slice(properties, func(i, j int) bool {
	//	p, q := properties[i], properties[j]
	//	return p[0] < q[0] || p[0] == q[0] && p[1] > q[1]
	//})
	//var st []int
	//for _, p := range properties {
	//	for len(st) > 0 && st[len(st)-1] < p[1] {
	//		st = st[:len(st)-1]
	//		ans++
	//	}
	//	st = append(st, p[1])
	//}
	//return
}
