/*
-------------------------------------
# @Time    : 2022/3/26 0:05:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 682_棒球比赛.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	ops := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(ops))
}

func calPoints(ops []string) (ans int) {
	var scores []int
	for i := 0; i < len(ops); i++ {
		if ops[i] == "C" {
			scores = scores[:len(scores)-1]
		} else if ops[i] == "D" {
			scores = append(scores, 2*scores[len(scores)-1])
		} else if ops[i] == "+" {
			scores = append(scores, scores[len(scores)-1]+scores[len(scores)-2])
		} else {
			v, _ := strconv.Atoi(ops[i])
			scores = append(scores, v)
		}
	}
	for _, score := range scores {
		ans += score
	}
	return
}
