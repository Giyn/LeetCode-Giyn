/*
-------------------------------------
# @Time    : 2022/2/22 10:51:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 718_最长重复子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(nums1, nums2))
}

func findLength(nums1 []int, nums2 []int) (ans int) {
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return
}
