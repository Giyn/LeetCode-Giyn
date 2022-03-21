/*
-------------------------------------
# @Time    : 2022/3/20 20:31:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : min_test.go
# @Software: GoLand
-------------------------------------
*/

package math

import (
	"fmt"
	"testing"
)

func TestMin(t *testing.T) {
	fmt.Println(Min([]int{1, 2, 3, 4, 6}))
	fmt.Println(Min(1, 2, 3, 4, 6))
}
