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
	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms)
	fmt.Println(ms.Sum("ap"))
}

type MapSum map[string]int

func Constructor() MapSum {
	return MapSum{}
}

func (ms MapSum) Insert(key string, val int) {
	ms[key] = val
}

func (ms MapSum) Sum(prefix string) (sum int) {
	for k, v := range ms {
		if strings.HasPrefix(k, prefix) {
			sum += v
		}
	}
	return
}
