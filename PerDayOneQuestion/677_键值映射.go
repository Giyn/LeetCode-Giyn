/*
-------------------------------------
# @Time    : 2021/11/14 0:48:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 677_键值映射.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	ms := Constructor677()
	ms.Insert("apple", 3)
	fmt.Println(ms)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms)
	fmt.Println(ms.Sum("ap"))
}

type MapSum map[string]int

func Constructor677() MapSum {
	return MapSum{}
}

func (this MapSum) Insert(key string, val int) {
	this[key] = val
}

func (this MapSum) Sum(prefix string) (sum int) {
	for k, v := range this {
		if strings.HasPrefix(k, prefix) {
			sum += v
		}
	}
	return
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
