/*
-------------------------------------
# @Time    : 2022/4/26 18:33:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 168_Excel表列名称.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	columnNumber := 701
	fmt.Println(convertToTitle(columnNumber))
}

func convertToTitle(columnNumber int) string {
	var ans []byte
	for columnNumber > 0 {
		columnNumber--
		ans = append(ans, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}
