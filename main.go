package main

import (
	"fmt"
	"sync"

	w "github.com/salawhaaat/bcwallet/wallet"
)

func main() {
	wallet := w.New(w.Bitcoin(100))
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			err := wallet.Withdraw(w.Bitcoin(5))
			if err != nil {
				panic(err)
			}
		}()
		go func() {
			defer wg.Done()
			err := wallet.Deposit(w.Bitcoin(10))
			if err != nil {
				panic(err)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Current balance:", wallet.Balance())
}
