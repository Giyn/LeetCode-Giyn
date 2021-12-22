/*
-------------------------------------
# @Time    : 2021/12/23 1:04:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1044_最长重复子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"index/suffixarray"
	"reflect"
	"unsafe"
)

func main() {
	s := "banana"
	fmt.Println(longestDupSubstring(s))
}

func longestDupSubstring(s string) string {
	// 求出后缀数组和高度数组
	n := len(s)
	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	// 高度数组中的最大值对应的就是最长的重复子串
	idx, maxH := 0, 0
	for i, h := range height {
		if h > maxH {
			idx, maxH = i, h
		}
	}
	return s[sa[idx] : int(sa[idx])+maxH]
}
