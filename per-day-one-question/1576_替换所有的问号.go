/*
-------------------------------------
# @Time    : 2022/1/5 12:38:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1576_替换所有的问号.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "??yw?ipkj?"
	fmt.Println(modifyString(s))
}

func modifyString(s string) string {
	ans := []byte(s)
	n := len(ans)
	for i, ch := range ans {
		if ch == '?' {
			for b := byte('a'); b <= 'c'; b++ {
				if !(i > 0 && ans[i-1] == b || i < n-1 && ans[i+1] == b) {
					ans[i] = b
					break
				}
			}
		}
	}
	return string(ans)
}
