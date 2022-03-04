/*
-------------------------------------
# @Time    : 2022/3/4 11:20:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 167_两数之和 II - 输入有序数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) (ans []int) {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			ans = []int{left + 1, right + 1}
			return
		}
	}
	return
}
