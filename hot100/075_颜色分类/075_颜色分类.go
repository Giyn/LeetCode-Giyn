/*
-------------------------------------
# @Time    : 2022/3/29 23:15:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 075_颜色分类.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	num0, num1, num2 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[num2] = 2
			nums[num1] = 1
			nums[num0] = 0
			num2++
			num1++
			num0++
		} else if nums[i] == 1 {
			nums[num2] = 2
			nums[num1] = 1
			num2++
			num1++
		} else {
			nums[num2] = 2
			num2++
		}
	}
}
