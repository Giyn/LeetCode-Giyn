/*
-------------------------------------
# @Time    : 2022/4/6 20:13:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_011_旋转数组的最小数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numbers := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray(numbers))
}

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := left + (right-left)>>1
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right-- // 无法确定在哪个范围，故缩小空间
		}
	}
	return numbers[left]
}
