/*
-------------------------------------
# @Time    : 2022/5/7 23:25:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 433_最小基因变化.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(minMutation(start, end, bank))
}

func minMutation(start string, end string, bank []string) int {
	const CHARS = "ACGT"
	bankSet, queue := map[string]bool{}, []string{start}
	for _, s := range bank {
		bankSet[s] = true
	}
	for step := 0; len(queue) > 0; step++ {
		for l := len(queue); l > 0; l-- {
			cur := queue[0]
			queue = queue[1:]
			for i, r := range cur {
				for _, nr := range CHARS {
					if r != nr {
						nxt := cur[:i] + string(nr) + cur[i+1:]
						if bankSet[nxt] {
							if nxt == end {
								return step + 1
							}
							bankSet[nxt] = false
							queue = append(queue, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}
