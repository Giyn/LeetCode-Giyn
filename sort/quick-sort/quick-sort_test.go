/*
-------------------------------------
# @Time    : 2022/3/21 20:20:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : quick-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package quick_sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
