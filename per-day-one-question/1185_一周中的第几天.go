/*
-------------------------------------
# @Time    : 2022/1/3 23:18:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1185_一周中的第几天.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	day := 31
	month := 8
	year := 2019
	fmt.Println(dayOfTheWeek(day, month, year))
}

func dayOfTheWeek(day int, month int, year int) string {
	var week = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var ans = 4
	for i := 1971; i < year; i++ {
		if (i%4 == 0 && i%100 != 0) || i%400 == 0 {
			ans++
		}
		ans += 365
	}
	for i := 1; i < month; i++ {
		if i == 2 && ((year%4 == 0 && year%100 != 0) || year%400 == 0) {
			ans++
		}
		ans += days[i-1]
	}
	ans += day
	return week[ans%7]
}
