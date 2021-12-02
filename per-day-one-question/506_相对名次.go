/*
-------------------------------------
# @Time    : 2021/12/2 0:58:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 506_相对名次.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))
}

var medal = [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

func findRelativeRanks(score []int) (ans []string) {
	n := len(score)
	arr := make([]scorePair, n)
	ans = make([]string, n)
	for i, s := range score {
		arr[i] = scorePair{i, s}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].score > arr[j].score })
	for i, p := range arr {
		if i < 3 {
			ans[p.index] = medal[i]
		} else {
			ans[p.index] = strconv.Itoa(i + 1)
		}
	}
	return
}

type scorePair struct {
	index int
	score int
}
