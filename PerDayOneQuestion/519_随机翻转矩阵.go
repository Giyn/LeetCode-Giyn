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

type Solution519 struct {
	m, n, total int
	mp          map[int]int
}

func main() {
	obj := Constructor519(3, 1)
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	obj.Reset()
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
	fmt.Println(obj.Flip())
}

func Constructor519(m int, n int) Solution519 {
	return Solution519{m, n, m * n, map[int]int{}}
}

func (this *Solution519) Flip() (ans []int) {
	x := rand.Intn(this.total)
	this.total--
	if y, ok := this.mp[x]; ok {
		ans = []int{y / this.n, y % this.n}
	} else {
		ans = []int{x / this.n, x % this.n}
	}
	if y, ok := this.mp[this.total]; ok {
		this.mp[x] = y
	} else {
		this.mp[x] = this.total
	}
	return
}

func (this *Solution519) Reset() {
	this.total = this.m * this.n
	this.mp = map[int]int{}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
