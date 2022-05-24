/*
-------------------------------------
# @Time    : 2021/11/27 1:00:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 519_随机翻转矩阵.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor(3, 1)
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	obj.Reset()
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
}

type Solution struct {
	m, n, total int
	mp          map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m * n, map[int]int{}}
}

func (s *Solution) Flip() (ans []int) {
	x := rand.Intn(s.total)
	s.total--
	if y, ok := s.mp[x]; ok {
		ans = []int{y / s.n, y % s.n}
	} else {
		ans = []int{x / s.n, x % s.n}
	}
	if y, ok := s.mp[s.total]; ok {
		s.mp[x] = y
	} else {
		s.mp[x] = s.total
	}
	return
}

func (s *Solution) Reset() {
	s.total = s.m * s.n
	s.mp = map[int]int{}
}
