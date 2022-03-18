/*
-------------------------------------
# @Time    : 2022/3/18 9:04:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2043_简易银行系统.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	b := Constructor2043([]int64{10, 100, 20, 50, 30})
	fmt.Println(b.Withdraw(3, 10))
	fmt.Println(b.Transfer(5, 1, 20))
	fmt.Println(b.Deposit(5, 20))
	fmt.Println(b.Transfer(3, 4, 15))
	fmt.Println(b.Withdraw(10, 50))
}

type Bank struct {
	balance []int64
}

func Constructor2043(balance []int64) Bank {
	return Bank{balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1-1 < len(this.balance) && account2-1 < len(this.balance) && this.balance[account1-1] >= money {
		this.balance[account1-1] -= money
		this.balance[account2-1] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account-1 < len(this.balance) {
		this.balance[account-1] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account-1 < len(this.balance) && this.balance[account-1] >= money {
		this.balance[account-1] -= money
		return true
	}
	return false
}
