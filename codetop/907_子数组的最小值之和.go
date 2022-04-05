/*
-------------------------------------
# @Time    : 2022/4/5 15:17:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 907_子数组的最小值之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(sumSubarrayMins(arr))
}

func sumSubarrayMins(arr []int) (ans int) {
	arr = append(arr, 0)
	stack := []int{-1}
	for i := 0; i < len(arr); i++ {
		for stack[len(stack)-1] != -1 && arr[i] <= arr[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 找到数组中的每一个数作为最小值的范围
			ans += arr[idx] * (i - idx) * (idx - stack[len(stack)-1]) % (1e9 + 7)
			ans %= 1e9 + 7
		}
		stack = append(stack, i)
	}
	return
}
