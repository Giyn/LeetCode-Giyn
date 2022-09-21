/*
-------------------------------------
# @Time    : 2022/9/21 23:47:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_066_构建乘积数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(a))
}

// O(1)
func constructArr(a []int) []int {
	length := len(a)
	ans := make([]int, length)
	if length == 0 {
		return ans
	}
	// ans[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 ans[0] = 1
	ans[0] = 1
	for i := 1; i < length; i++ {
		ans[i] = a[i-1] * ans[i-1]
	}

	// R 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	R := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 ans[i]，右边的乘积为 R
		ans[i] = ans[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= a[i]
	}
	return ans
}

// O(n)
//func constructArr(a []int) []int {
//	length := len(a)
//
//	// L 和 R 分别表示左右两侧的乘积列表
//	L, R, ans := make([]int, length), make([]int, length), make([]int, length)
//
//	// L[i] 为索引 i 左侧所有元素的乘积
//	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
//	L[0] = 1
//	for i := 1; i < length; i++ {
//		L[i] = a[i-1] * L[i-1]
//	}
//
//	// R[i] 为索引 i 右侧所有元素的乘积
//	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
//	R[length-1] = 1
//	for i := length - 2; i >= 0; i-- {
//		R[i] = a[i+1] * R[i+1]
//	}
//
//	// 对于索引 i，除 a[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
//	for i := 0; i < length; i++ {
//		ans[i] = L[i] * R[i]
//	}
//	return ans
//}
