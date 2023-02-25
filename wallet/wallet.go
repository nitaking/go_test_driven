package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() interface{} {
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	return w.balance
}
