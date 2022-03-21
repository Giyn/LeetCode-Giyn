/*
-------------------------------------
# @Time    : 2022/3/21 21:39:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : shell-sort.go
# @Software: GoLand
-------------------------------------
*/

package shell_sort

func ShellSort(nums []int) {
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			j := i
			cur := nums[i]
			for j-gap >= 0 && nums[j-gap] > cur {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = cur
		}
	}
}
