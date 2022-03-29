/*
-------------------------------------
# @Time    : 2022/3/29 9:04:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2024_考试的最大困扰度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	answerKey := "TTFF"
	k := 2
	fmt.Println(maxConsecutiveAnswers(answerKey, k))
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	maxCountKey := func(c byte) (ans int) {
		left, cnt := 0, 0
		for right := range answerKey {
			if answerKey[right] != c {
				cnt++
			}
			for cnt > k {
				if answerKey[left] != c {
					cnt--
				}
				left++
			}
			ans = Max(ans, right-left+1)
		}
		return
	}
	return Max(maxCountKey('T'), maxCountKey('F'))
}
