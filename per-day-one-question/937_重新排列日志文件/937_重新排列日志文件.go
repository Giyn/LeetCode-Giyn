/*
-------------------------------------
# @Time    : 2022/5/3 0:56:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 937_重新排列日志文件.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	fmt.Println(reorderLogFiles(logs))
}

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		logi, logj := strings.SplitN(logs[i], " ", 2), strings.SplitN(logs[j], " ", 2)
		if unicode.IsLetter(rune(logi[1][0])) && unicode.IsDigit(rune(logj[1][0])) {
			return true
		}
		if unicode.IsLetter(rune(logi[1][0])) && unicode.IsLetter(rune(logj[1][0])) {
			if logi[1] != logj[1] {
				return logi[1] < logj[1]
			}
			return logi[0] < logj[0]
		}
		return false
	})
	return logs
}
