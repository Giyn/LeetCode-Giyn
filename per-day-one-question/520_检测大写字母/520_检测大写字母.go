/*
-------------------------------------
# @Time    : 2021/11/13 0:01:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 520_检测大写字母.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	word := "a"
	fmt.Println(detectCapitalUse(word))
}

func detectCapitalUse(word string) bool {
	// 若第 1 个字母为小写，则需额外判断第 2 个字母是否为小写
	if len(word) >= 2 && unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}

	// 无论第 1 个字母是否大写，其他字母必须与第 2 个字母的大小写相同
	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
	}
	return true
}

//func detectCapitalUse(word string) bool {
//	if len(word) == 1 {
//		return true
//	}
//	flag := -1 // 0:全大写 1:全小写 2:仅首字母大写
//	wr := []rune(word)
//	if wr[0]-'A' <= 25 && wr[1]-'A' <= 25 {
//		flag = 0
//	} else if wr[0]-'A' >= 32 && wr[1]-'A' >= 32 {
//		flag = 1
//	} else if wr[0]-'A' <= 25 && wr[1]-'A' >= 32 {
//		flag = 2
//	}
//	switch flag {
//	case 0:
//		for i := 2; i < len(wr); i++ {
//			if wr[i]-'A' >= 32 {
//				return false
//			}
//		}
//	case 1:
//		for i := 2; i < len(wr); i++ {
//			if wr[i]-'A' <= 25 {
//				return false
//			}
//		}
//	case 2:
//		for i := 2; i < len(wr); i++ {
//			if wr[i]-'A' <= 25 {
//				return false
//			}
//		}
//	default:
//		return false
//	}
//	return true
//}
