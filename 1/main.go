package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if amount <= b.balance {
		b.balance -= amount
	} else {
		fmt.Println("There are not enough funds in the account.")
	}
}

func main() {
	account := BankAccount{balance: 0}

	var wg sync.WaitGroup

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Deposit(10)
		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.Withdraw(500)
		}()
	}

	wg.Wait()

	fmt.Printf("Final Balance: %d\n", account.balance)
}
