/*
-------------------------------------
# @Time    : 2022/3/22 11:11:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : counting-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package counting_sort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	CountingSort(nums)
	fmt.Println(nums)
}
