/*
-------------------------------------
# @Time    : 2022/3/9 10:02:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 724_寻找数组的中心下标.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		if 2*sum+num == total {
			return i
		}
		sum += num
	}
	return -1
}
