/*
-------------------------------------
# @Time    : 2022/2/27 0:14:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 553_最优除法.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{1000, 100, 10, 2}
	fmt.Println(optimalDivision(nums))
}

func optimalDivision(nums []int) string {
	var ans []string
	for _, num := range nums {
		ans = append(ans, strconv.Itoa(num))
	}
	if len(ans) < 3 {
		return strings.Join(ans, "/")
	}
	return ans[0] + "/(" + strings.Join(ans[1:], "/") + ")"
}
