/*
-------------------------------------
# @Time    : 2021/12/17 9:41:54
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 134_加油站.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	var start int
	var curSum int
	var totalSum int

	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}
