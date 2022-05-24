/*
-------------------------------------
# @Time    : 2021/12/11 11:09:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 911_在线选举.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	obj := Constructor(persons, times)
	fmt.Println(obj.Q(3))
	fmt.Println(obj.Q(12))
	fmt.Println(obj.Q(25))
	fmt.Println(obj.Q(15))
	fmt.Println(obj.Q(24))
	fmt.Println(obj.Q(8))
}

type TopVotedCandidate struct {
	ans, times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(times)
	ans, cnts, cur := make([]int, n), make([]int, n), -1
	for i := range times {
		cnts[persons[i]]++
		if cur == -1 || cnts[persons[i]] >= cnts[cur] {
			cur = persons[i]
		}
		ans[i] = cur
	}
	return TopVotedCandidate{ans, times}
}

func (tvc *TopVotedCandidate) Q(t int) int {
	left, right := 0, len(tvc.times)
	for left < right {
		mid := left + (right-left)>>1
		if t >= tvc.times[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return tvc.ans[left-1]
}
