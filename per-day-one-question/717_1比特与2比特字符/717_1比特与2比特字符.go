/*
-------------------------------------
# @Time    : 2022/2/20 1:00:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 717_1比特与2比特字符.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	bits := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(bits))
}

func isOneBitCharacter(bits []int) (ans bool) {
	for i := 0; i < len(bits); {
		if bits[i] == 0 {
			ans = true
			i++
		} else {
			ans = false
			i += 2
		}
	}
	return
}
