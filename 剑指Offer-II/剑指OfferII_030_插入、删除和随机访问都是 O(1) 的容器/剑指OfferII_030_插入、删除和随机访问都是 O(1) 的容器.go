/*
-------------------------------------
# @Time    : 2022/4/20 13:05:51
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_030_插入、删除和随机访问都是 O(1) 的容器.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := Constructor()
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

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.mp[val]; ok {
		return false
	}
	rs.mp[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, ok := rs.mp[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1
	rs.nums[idx] = rs.nums[last]
	rs.mp[rs.nums[idx]] = idx
	rs.nums = rs.nums[:last]
	delete(rs.mp, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}
