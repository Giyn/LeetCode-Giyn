/*
-------------------------------------
# @Time    : 2021/11/10 0:03:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 495_提莫攻击.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	timeSeries := []int{1, 2}
	duration := 2
	fmt.Println(findPoisonedDuration(timeSeries, duration))
}

func findPoisonedDuration(timeSeries []int, duration int) (ans int) {
	for i := 1; i < len(timeSeries); i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if diff >= duration {
			ans += duration
		} else {
			ans += diff
		}
	}
	return ans + duration
}
