/*
-------------------------------------
# @Time    : 2021/10/29 15:35:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 335_路径交叉.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	distance := []int{1, 2, 3, 4}
	fmt.Println(isSelfCrossing(distance))
}

func isSelfCrossing(distance []int) bool {
	if len(distance) < 4 {
		return false
	}
	for i := 3; i < len(distance); i++ {
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		if i >= 4 && distance[i-1] == distance[i-3] && distance[i]+distance[i-4] >= distance[i-2] {
			return true
		}
		if i >= 5 && distance[i-1] <= distance[i-3] && distance[i-2] > distance[i-4] && distance[i]+distance[i-4] >= distance[i-2] && distance[i-1]+distance[i-5] >= distance[i-3] {
			return true
		}
	}
	return false
}
