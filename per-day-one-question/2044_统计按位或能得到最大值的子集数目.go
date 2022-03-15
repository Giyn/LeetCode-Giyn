/*
-------------------------------------
# @Time    : 2022/3/15 10:08:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2044_统计按位或能得到最大值的子集数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 5}
	fmt.Println(countMaxOrSubsets(nums))
}

func countMaxOrSubsets(nums []int) (ans int) {
	// 回溯
	maxOr := 0
	var dfs func(pos, or int)
	dfs = func(pos, or int) {
		if pos == len(nums) {
			if or > maxOr {
				maxOr = or
				ans = 1
			} else if or == maxOr {
				ans++
			}
			return
		}
		dfs(pos+1, or|nums[pos])
		dfs(pos+1, or)
	}
	dfs(0, 0)
	return

	// 枚举
	//maxOr := 0
	//for i := 1; i < 1<<len(nums); i++ {
	//	or := 0
	//	for j, num := range nums {
	//		if i>>j&1 == 1 {
	//			or |= num
	//		}
	//	}
	//	if or > maxOr {
	//		maxOr = or
	//		ans = 1
	//	} else if or == maxOr {
	//		ans++
	//	}
	//}
	//return
}
