/*
-------------------------------------
# @Time    : 2022/1/9 23:14:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1629_按键持续时间最长的键.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	releaseTimes := []int{9, 29, 49, 50}
	keysPressed := "cbcd"
	fmt.Println(slowestKey(releaseTimes, keysPressed))
}

func slowestKey(releaseTimes []int, keysPressed string) (ans byte) {
	maxTime := releaseTimes[0]
	ans = keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		key := keysPressed[i]
		time := releaseTimes[i] - releaseTimes[i-1]
		if time > maxTime || time == maxTime && key > ans {
			maxTime = releaseTimes[i] - releaseTimes[i-1]
			ans = keysPressed[i]
		}
	}
	return ans
}
