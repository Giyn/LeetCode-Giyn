/*
-------------------------------------
# @Time    : 2021/12/18 19:23:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 860_柠檬水找零.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	bills := []int{5, 5, 10}
	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	if bills[0] != 5 {
		return false
	}
	mp := make(map[int]int, 3)
	for _, bill := range bills {
		mp[bill]++
		if bill == 5 {
			continue
		} else if bill == 10 {
			if mp[5] > 0 {
				mp[5]--
			} else {
				return false
			}
		} else if bill == 20 {
			if mp[10] > 0 {
				mp[10]--
				if mp[5] > 0 {
					mp[5]--
				} else {
					return false
				}
			} else {
				if mp[5] > 2 {
					mp[5] -= 3
				} else {
					return false
				}
			}
		}
	}
	return true
}
