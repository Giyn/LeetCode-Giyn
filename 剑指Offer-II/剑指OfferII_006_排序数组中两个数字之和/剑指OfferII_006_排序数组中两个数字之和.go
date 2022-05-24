/*
-------------------------------------
# @Time    : 2022/3/4 10:56:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_006_排序数组中两个数字之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4, 6, 10}
	target := 8
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) (ans []int) {
	// 双指针
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			ans = []int{left, right}
			return
		}
	}
	return

	// 哈希表
	//mp := make(map[int]int, len(numbers))
	//for i, num := range numbers {
	//	mp[num] = i
	//}
	//for i, num := range numbers {
	//	if v, ok := mp[target-num]; ok {
	//		ans = []int{i, v}
	//		return
	//	}
	//}
	//return
}
