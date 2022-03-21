/*
-------------------------------------
# @Time    : 2022/3/21 21:40:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : shell-sort_test.go
# @Software: GoLand
-------------------------------------
*/

package shell_sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	nums := []int{6, 2, 7, 5, 9, 27, 12, 52, 57, 42}
	ShellSort(nums)
	fmt.Println(nums)
}
