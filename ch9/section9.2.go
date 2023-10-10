package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

var balance int

func deposit(amount int) { balance += amount }

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()

	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}

	return true
}

func main() {
	fmt.Println(Withdraw(10))
	Deposit(300)
	fmt.Println("Deposit(300): Balance: ", Balance())
	fmt.Println(Withdraw(10))
}
