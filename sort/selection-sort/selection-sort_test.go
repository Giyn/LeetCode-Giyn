/*
-------------------------------------
# @Time    : 2022/3/21 17:12:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : selection-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package selection_sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	SelectionSort(nums)
	fmt.Println(nums)
}
