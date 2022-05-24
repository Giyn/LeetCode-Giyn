/*
-------------------------------------
# @Time    : 2022/1/4 23:46:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2022_将一维数组转变成二维数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4}
	m := 2
	n := 2
	fmt.Println(construct2DArray(original, m, n))
}

func construct2DArray(original []int, m int, n int) (ans [][]int) {
	length := len(original)
	if length != m*n {
		return
	}
	for i := 0; i < length; i += n {
		ans = append(ans, original[i:i+n])
	}
	return
}
