/*
-------------------------------------
# @Time    : 2022/3/13 0:48:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 393_UTF-8 编码验证.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	data := []int{250, 145, 145, 145, 145}
	fmt.Println(validUtf8(data))
}

func validUtf8(data []int) bool {
	const one = 1 << 7
	const two = one + (1 << 6)
	for i := 0; i < len(data); {
		count := 0 // 1的个数
		for count < 8 && (data[i]>>(7-count))&1 == 1 {
			count++
		}
		if count == 0 {
			i++
			continue
		} else if count == 1 || count > 4 || i+count-1 >= len(data) {
			return false
		}
		for j := 1; j < count; j++ {
			if data[i+j]&two != one {
				return false
			}
		}
		i += count
	}
	return true
}
