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
	stockPrice := Constructor2034()
	stockPrice.Update(1, 10)
	stockPrice.Update(2, 5)
	fmt.Println(stockPrice.Current())
	fmt.Println(stockPrice.Maximum())
	stockPrice.Update(1, 3)
	fmt.Println(stockPrice.Maximum())
	stockPrice.Update(4, 2)
	fmt.Println(stockPrice.Minimum())
}

func Constructor2034() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if prevPrice := this.timePriceMap[timestamp]; prevPrice > 0 {
		if times, _ := this.prices.Get(prevPrice); times.(int) > 1 {
			this.prices.Put(prevPrice, times.(int)-1)
		} else {
			this.prices.Remove(prevPrice)
		}
	}
	times := 0
	if val, ok := this.prices.Get(price); ok {
		times = val.(int)
	}
	this.prices.Put(price, times+1)
	this.timePriceMap[timestamp] = price
	if timestamp > this.maxTimestamp {
		this.maxTimestamp = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.timePriceMap[this.maxTimestamp]
}

func (this *StockPrice) Maximum() int {
	return this.prices.Right().Key.(int)
}

func (this *StockPrice) Minimum() int {
	return this.prices.Left().Key.(int)
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
