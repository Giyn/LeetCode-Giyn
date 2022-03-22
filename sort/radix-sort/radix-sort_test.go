/*
-------------------------------------
# @Time    : 2022/3/22 11:57:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : radix-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package radix_sort

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	RadixSort(nums)
	fmt.Println(nums)
}
