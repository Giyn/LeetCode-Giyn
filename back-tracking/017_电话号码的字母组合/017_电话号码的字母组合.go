/*
-------------------------------------
# @Time    : 2021/12/3 15:53:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 017_电话号码的字母组合.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) (ans []string) {
	if digits == "" {
		return nil
	}
	var path []byte
	letterMap := [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(digits) {
			ans = append(ans, string(path))
			return
		}
		digit := digits[index] - '0'
		letters := letterMap[digit]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
