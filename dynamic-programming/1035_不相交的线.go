/*
-------------------------------------
# @Time    : 2022/2/22 11:13:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1035_不相交的线.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(nums1)][len(nums2)]
}
