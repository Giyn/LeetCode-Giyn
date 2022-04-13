/*
-------------------------------------
# @Time    : 2022/4/13 1:24:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 380_O(1) 时间插入、删除和获取随机元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := Constructor380()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(2))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())
}

type RandomizedSet struct {
	nums []int
	mp   map[int]int
}

func Constructor380() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}
	this.mp[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.mp[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[idx] = this.nums[last]
	this.mp[this.nums[idx]] = idx
	this.nums = this.nums[:last]
	delete(this.mp, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
