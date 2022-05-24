/*
-------------------------------------
# @Time    : 2022/1/23 13:32:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2034_股票价格波动.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

type StockPrice struct {
	prices       *redblacktree.Tree
	timePriceMap map[int]int
	maxTimestamp int
}

func main() {
	stockPrice := Constructor()
	stockPrice.Update(1, 10)
	stockPrice.Update(2, 5)
	fmt.Println(stockPrice.Current())
	fmt.Println(stockPrice.Maximum())
	stockPrice.Update(1, 3)
	fmt.Println(stockPrice.Maximum())
	stockPrice.Update(4, 2)
	fmt.Println(stockPrice.Minimum())
}

func Constructor() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0}
}

func (sp *StockPrice) Update(timestamp int, price int) {
	if prevPrice := sp.timePriceMap[timestamp]; prevPrice > 0 {
		if times, _ := sp.prices.Get(prevPrice); times.(int) > 1 {
			sp.prices.Put(prevPrice, times.(int)-1)
		} else {
			sp.prices.Remove(prevPrice)
		}
	}
	times := 0
	if val, ok := sp.prices.Get(price); ok {
		times = val.(int)
	}
	sp.prices.Put(price, times+1)
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}

func (sp *StockPrice) Maximum() int {
	return sp.prices.Right().Key.(int)
}

func (sp *StockPrice) Minimum() int {
	return sp.prices.Left().Key.(int)
}
