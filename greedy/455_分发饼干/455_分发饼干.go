/*
-------------------------------------
# @Time    : 2021/12/12 4:14:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 455_分发饼干.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) (ans int) {
	sort.Ints(g)
	sort.Ints(s)
	index := 0
	for i := 0; i < len(s); i++ {
		if index < len(g) && s[i] >= g[index] {
			ans++
			index++
			continue
		}
	}
	return
}
