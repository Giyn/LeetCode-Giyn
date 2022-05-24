/*
-------------------------------------
# @Time    : 2022/4/9 14:15:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 780_到达终点.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	sx := 1
	sy := 1
	tx := 3
	ty := 5
	fmt.Println(reachingPoints(sx, sy, tx, ty))
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > sx && ty > sy {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	return (tx == sx && ty >= sy && (ty-sy)%sx == 0) || (ty == sy && tx >= sx && (tx-sx)%sy == 0)
}
