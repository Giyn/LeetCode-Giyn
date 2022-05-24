/*
-------------------------------------
# @Time    : 2021/12/21 0:30:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1154_一年中的第几天.go
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
	date := "2003-03-01"
	fmt.Println(dayOfYear(date))
}

func dayOfYear(date string) (ans int) {
	var days = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var isRun func(int) bool
	isRun = func(year int) bool {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return true
		} else {
			return false
		}
	}
	dateSlice := strings.Split(date, "-")
	year, _ := strconv.Atoi(dateSlice[0])
	month, _ := strconv.Atoi(dateSlice[1])
	day, _ := strconv.Atoi(dateSlice[2])
	if isRun(year) {
		days[1] = 29
	}
	for i := 0; i < month-1; i++ {
		ans += days[i]
	}
	return ans + day
}
