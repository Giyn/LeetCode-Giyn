/*
-------------------------------------
# @Time    : 2022/4/20 11:14:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 388_文件的最长绝对路径.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	input := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	fmt.Println(lengthLongestPath(input))
}

func lengthLongestPath(input string) (ans int) {
	st := []int{}
	for i, n := 0, len(input); i < n; {
		// 检测当前文件的深度
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}

		// 统计当前文件名的长度
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++ // 跳过换行符

		for len(st) >= depth {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			length += st[len(st)-1] + 1
		}
		if isFile {
			ans = Max(ans, length)
		} else {
			st = append(st, length)
		}
	}
	return
}
