/*
-------------------------------------
# @Time    : 2022/5/23 1:49:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 169_多数元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}

func majorityElement(nums []int) (ans int) {
	count := 0
	for _, num := range nums {
		if count == 0 {
			ans = num
			count++
		} else {
			if num == ans {
				count++
			} else {
				count--
			}
		}
	}
	return
}
