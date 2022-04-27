/*
-------------------------------------
# @Time    : 2022/4/26 21:24:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 398_随机数索引.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	solution := Constructor398([]int{1, 2, 3, 3, 3})
	fmt.Println(solution.Pick(3))
	fmt.Println(solution.Pick(1))
	fmt.Println(solution.Pick(3))
}

type Solution398 struct {
	mp map[int][]int
}

func Constructor398(nums []int) Solution398 {
	mp := map[int][]int{}
	for i, num := range nums {
		mp[num] = append(mp[num], i)
	}
	return Solution398{mp}
}

func (this *Solution398) Pick(target int) int {
	s := this.mp[target]
	return s[rand.Intn(len(s))]
}
