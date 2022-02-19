/*
-------------------------------------
# @Time    : 2022/2/19 20:34:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 969_煎饼排序.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 4, 1}
	fmt.Println(pancakeSort(arr))
}

func pancakeSort(arr []int) (ans []int) {
	reverse := func(arr *[]int, idx int) {
		for i := 0; i < (idx+1)/2; i++ {
			(*arr)[i], (*arr)[idx-i] = (*arr)[idx-i], (*arr)[i]
		}
	}
	for n := len(arr); n > 1; n-- {
		idx := 0
		for i, v := range arr[:n] {
			if v > arr[idx] {
				idx = i
			}
		}
		if idx == n-1 {
			continue
		}
		reverse(&arr, idx)
		ans = append(ans, idx+1)
		reverse(&arr, n-1)
		ans = append(ans, n)
	}
	return
}
