/*
-------------------------------------
# @Time    : 2021/10/14 14:56:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_II_069_山峰数组的顶部.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	arr := []int{0, 1, 0}
	fmt.Println(peakIndexInMountainArray(arr))
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
	for left <= right {
		mid := left + (right - left) >> 1
		if arr[mid-1] < arr[mid] && arr[mid+1] < arr[mid] {
			return mid
		} else if arr[mid-1] < arr[mid] {
			left = mid + 1
		} else {
			right = mid -1
		}
	}
	return right
}
