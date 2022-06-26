/*
-------------------------------------
# @Time    : 2022/6/27 3:39:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 071_简化路径.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/a/./b/../../c/"
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	var stack []string
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	return "/" + strings.Join(stack, "/")
}
