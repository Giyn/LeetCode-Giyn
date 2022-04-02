/*
-------------------------------------
# @Time    : 2022/4/3 0:00:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 744_寻找比目标字母大的最小字母.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')
	fmt.Println(nextGreatestLetter(letters, target))
}

func nextGreatestLetter(letters []byte, target byte) (ans byte) {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)>>1
		if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return letters[left]
}
