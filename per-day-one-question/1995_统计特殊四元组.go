/*
-------------------------------------
# @Time    : 2021/12/30 10:49:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1995_统计特殊四元组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 5}
	fmt.Println(countQuadruplets(nums))
}

func countQuadruplets(nums []int) (ans int) {
	cnt := map[int]int{}
	for b := len(nums) - 3; b >= 1; b-- {
		for _, x := range nums[b+2:] {
			cnt[x-nums[b+1]]++
		}
		for _, x := range nums[:b] {
			if sum := x + nums[b]; cnt[sum] > 0 {
				ans += cnt[sum]
			}
		}
	}
	return
}
