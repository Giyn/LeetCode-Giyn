/*
-------------------------------------
# @Time    : 2022/2/18 0:24:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1791_找出星型图的中心节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	fmt.Println(findCenter(edges))
}

func findCenter(edges [][]int) (ans int) {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}
