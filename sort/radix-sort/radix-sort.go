/*
-------------------------------------
# @Time    : 2022/3/22 11:32:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : radix-sort.go
# @Software: GoLand
-------------------------------------
*/

package radix_sort

import . "LeetCodeGiyn/utils/math"

func RadixSort(nums []int) {
	maxVal := Max(nums)
	maxBit := 0 // 最大值的位数
	for maxVal > 0 {
		maxBit++
		maxVal /= 10
	}
	tmp := make([]int, len(nums))
	count := new([10]int)
	for i, radix := 0, 1; i < maxBit; i, radix = i+1, radix*10 {
		for j := 0; j < 10; j++ {
			count[j] = 0
		}
		for j := 0; j < len(nums); j++ {
			count[(nums[j]/radix)%10]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j-1]
		}
		for j := len(nums) - 1; j >= 0; j-- {
			tmp[count[(nums[j]/radix)%10]-1] = nums[j]
			count[(nums[j]/radix)%10]--
		}
		for j := 0; j < len(nums); j++ {
			nums[j] = tmp[j]
		}
	}
}
