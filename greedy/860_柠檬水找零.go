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
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
			continue
		} else if bill == 10 {
			ten++
			if five > 0 {
				five--
			} else {
				return false
			}
		} else if bill == 20 {
			if ten > 0 {
				ten--
				if five > 0 {
					five--
				} else {
					return false
				}
			} else {
				if five > 2 {
					five -= 3
				} else {
					return false
				}
			}
		}
	}
	return true
}
