/*
-------------------------------------
# @Time    : 2022/8/29 23:37:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 137_只出现一次的数字 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	// 对于数组中非答案的元素，每一个元素都出现了 3 次，对应着第 i 个二进制位的 3 个 0 或 3 个 1，无论是哪一种情况，它们的和都是 3 的倍数（即和为 0 或 3）
	// 因此答案的第 i 个二进制位就是数组中所有元素的第 i 个二进制位之和除以 3 的余数
	ans := int32(0)
	for i := 0; i < 32; i++ {
		bitNum := 0
		for _, num := range nums {
			bitNum += (num >> i) & 1
		}
		if bitNum%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
